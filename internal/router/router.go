package router

import (
	"bluebell/api"
	"bluebell/config"
	"bluebell/logger"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

//路由请求

func InitRouterAndServe(cfg *config.AppConf) {

	//初始化gin engine
	router := gin.New()
	router.Use(logger.GinLogger(), logger.GinRecovery(true)) //使用zap日志中间件

	//绑定路由
	router.POST("/signup", api.SignUpHandler)

	//启动服务
	if err := router.Run(":" + strconv.Itoa(cfg.Port)); err != nil {
		zap.L().Error("run server failed" + err.Error())
		fmt.Println("run server failed" + err.Error())
		return
	}
}
