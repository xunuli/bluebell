package middlewares

import (
	"bluebell/api"
	"bluebell/pkg/jwt"
	"github.com/gin-gonic/gin"
	"strings"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//客户端携带Token有三种方式：1、放在请求头 2、放在请求体 3、放在URI （2、3 不合适）
		//这里假设Token放在Header的Authorization中，并使用Bearer开头
		//Authorization: Bearer xxxxxxx.xxx.xxx  / X-TOKEN: xxx.xxx.xx
		//这里的具体实现方式要依据你的实际业务情况决定

		//从请求头中获取jwt
		autoHeader := c.Request.Header.Get("Authorization")
		if autoHeader == "" {
			//JWT如果为空，说明未登录
			api.ResponseError(c, api.CodeNeedLogin)
			c.Abort() //终止向后执行
			return
		}
		//按空格分隔两个部分
		parts := strings.SplitN(autoHeader, " ", 2)
		//校验JWT的格式是否正确
		if len(parts) != 2 && parts[0] != "Bearer" {
			api.ResponseError(c, api.CodeInvalidToken)
			c.Abort()
			return
		}
		//parts[1]是获取到的tokenString，我们使用之间定义好的JWT解析函数来解析
		mc, err := jwt.ParseToken(parts[1])
		if err != nil {
			api.ResponseError(c, api.CodeInvalidToken)
			c.Abort()
			return
		}
		//将当前请求的userid信息保存到请求的上下文c中
		c.Set("Username", mc.Username)
		//继续处理后续的中间件和处理函数，可以通过c.get()来获取当前请求的用户信息id
		c.Next()
	}
}
