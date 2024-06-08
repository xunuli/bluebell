package mysql

import (
	"bluebell/config"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var (
	db *gorm.DB
)

// 用于初始化mysql连接
func InitMysql(cfg *config.DbConf) {
	//dsn数据库连接配置
	dsn := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", cfg.User,
		cfg.Password, cfg.Host, cfg.Port, cfg.Dbname)
	zap.L().Info("mysql addr:" + dsn)
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect mysql, err:" + err.Error())
	}

	//转化为通用数据库接口
	sqldb, err := db.DB()
	if err != nil {
		panic("fetch db connection err:" + err.Error())
	}

	sqldb.SetMaxIdleConns(cfg.MaxIdleConn)                                      //设置最大空闲连接
	sqldb.SetMaxOpenConns(cfg.MaxOpenConn)                                      //最大打开连接数
	sqldb.SetConnMaxLifetime(time.Duration(cfg.MaxIdleTime * int(time.Second))) //设置空闲时间为(s)
}

func GetDB() *gorm.DB {
	return db
}
