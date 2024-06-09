package router

import (
	"bluebell/api"
	"bluebell/logger"
	"bluebell/middlewares"
	"github.com/gin-gonic/gin"
	"net/http"
)

//路由请求

func InitRouterAndServe() *gin.Engine {

	//初始化gin engine
	router := gin.New()
	router.Use(logger.GinLogger(), logger.GinRecovery(true)) //使用zap日志中间件

	//绑定路由
	//注册
	router.POST("/signup", api.SignUpHandler)
	//登录
	router.POST("/login", api.LoginHandler)

	//针对登录用户发送的请求处理
	//http是无状态的，无法记录处于登录的用户的状态，因此要采用JWT记录用户的状态
	//采用gin中间件对JWT进行校验
	router.GET("/ping", middlewares.JWTAuthMiddleware(), func(c *gin.Context) {
		//如果是登录的用户，判断请求头中是否有 有效的JWT？
		username, _ := c.Get("Username")
		c.JSON(http.StatusOK, gin.H{
			"msg": username.(string) + " loging...",
		})
	})

	return router
}
