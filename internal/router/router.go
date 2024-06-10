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
	//心跳检测
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	//V1版本，路由组
	v1 := router.Group("api/v1")
	//注册
	v1.POST("/signup", api.SignUpHandler)
	//登录
	v1.POST("/login", api.LoginHandler)
	//V1路由组使用JWT中间件
	//针对登录用户发送的请求处理
	//http是无状态的，无法记录处于登录的用户的状态，因此要采用JWT记录用户的状态
	//采用gin中间件对JWT进行校验
	v1.Use(middlewares.JWTAuthMiddleware())
	{
		v1.GET("/community", api.CommunityHandler)
		v1.GET("/community/:id", api.CommunityDetailHandler)
	}

	return router
}
