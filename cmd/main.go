package main

import (
	"bluebell/config"
	"bluebell/internal/dao/mysql"
	"bluebell/internal/dao/redis"
	"bluebell/internal/router"
	"bluebell/logger"
	"bluebell/pkg/snowflake"
	"fmt"
	"go.uber.org/zap"
	"strconv"
)

var conf *config.GlobalConfig

// 常用goweb开发
func Init() {
	//1. 加载配置
	conf = config.GetGlobalConf()
	//2. 初始化日志
	logger.InitLog(&conf.LogConfig, conf.AppConfig.RunMode)
	//3. 初始化MySQL连接
	mysql.InitMysql(&conf.DbConfig)
	//4. 初始化Redis连接
	redis.InitRedis(&conf.RedisConfig)
	//初始化snowflake生成id模块
	snowflake.Init(&conf.SfConfig)

	zap.L().Info("init success....")
}

func main() {
	Init()
	res := mysql.GetDB()
	fmt.Println(res.Name())
	//5. 注册路由
	r := router.InitRouterAndServe()
	//6. 启动服务（优雅关机）
	//启动服务
	if err := r.Run(":" + strconv.Itoa(conf.AppConfig.Port)); err != nil {
		zap.L().Error("run server failed" + err.Error())
		fmt.Println("run server failed" + err.Error())
		return
	}
}
