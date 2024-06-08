package api

import (
	"bluebell/internal/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

//请求进来，路由到对应的接口处理函数

// 用户注册
func SignUpHandler(c *gin.Context) {
	//1. 获取参数和参数校验
	req := &service.SignUpRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		//请求参数的格式有误，直接返回响应
		zap.L().Error("SignUp with invaild param", zap.Error(err)) //记录日志
		ResponseError(c, CodeInvalidParam)
		return
	}
	//2. 业务处理，调用服务
	service.SignUp()
	//3. 返回响应
	c.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})
}
