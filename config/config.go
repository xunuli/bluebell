package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"sync"
)

var (
	config GlobalConfig //全局服务配置结构，单例模式
	once   sync.Once    //只执行一次初始化
)

// GlobalConfig 全局服务配置结构
type GlobalConfig struct {
	AppConfig   AppConf   `yaml:"app" mapstructure:"app"`      //服务配置
	LogConfig   LogConf   `yaml:"log" mapstructure:"log"`      // 日志配置
	DbConfig    DbConf    `yaml:"db" mapstructure:"mysql"`     // db配置
	RedisConfig RedisConf `yaml:"redis" mapstructure:"redis"`  // redis配置
	SfConfig    SfConf    `yaml:"snowflake" mapstructure:"sf"` //雪花算法唯一ID生成配置
}

// 服务配置
type AppConf struct {
	AppName string `yaml:"app_name" mapstructure:"app_name"`
	Version string `yaml:"version" mapstructure:"version"`
	Port    int    `yaml:"port" mapstructure:"port"`
	RunMode string `yaml:"run_mode" mapstructure:"run_mode"`
}

// 日志配置
type LogConf struct {
	Lever      string `yaml:"lever" mapstructure:"level"`             //日志级别
	LogPattern string `yaml:"log_pattern" mapstructure:"log_pattern"` //日志模式
	LogPath    string `yaml:"log_path" mapstructure:"log_path"`       //日志路径
	Filename   string `yaml:"filename" mapstructure:"filename"`       //日志名
	MaxSize    int    `yaml:"max_size" mapstructure:"max_size"`       //日志文件
	MaxAge     int    `yaml:"max_age" mapstructure:"max_age"`         //日志保存天数
	MaxBackups int    `yaml:"max_backups" mapstructure:"max_backups"` //日志备份数量
}

// 数据库配置
type DbConf struct {
	Host        string `yaml:"host" mapstructure:"host"`                   //数据库地址
	Port        int    `yaml:"port" mapstructure:"port"`                   //端口
	User        string `yaml:"user" mapstructure:"user"`                   //用户
	Password    string `yaml:"password" mapstructure:"password"`           //密码
	Dbname      string `yaml:"dbname" mapstructure:"dbname"`               //数据库名字
	MaxIdleConn int    `yaml:"max_idle_conn" mapstructure:"max_idle_conn"` //最大空闲连接
	MaxOpenConn int    `yaml:"max_open_conn" mapstructure:"max_open_conn"` //最大连接数
	MaxIdleTime int    `yaml:"max_idle_time" mapstructure:"max_idle_time"` //连接最大空闲数量
}

// redis配置
type RedisConf struct {
	Host     string `yaml:"host" mapstructure:"host"`         //Redis地址
	Port     int    `yaml:"port" mapstructure:"port"`         //端口
	Db       int    `yaml:"db" mapstructure:"db"`             //数据库名0-15
	Password string `yaml:"password" mapstructure:"password"` //密码
	Poolsize int    `yaml:"poolsize" mapstructure:"poolsize"` //客户端连接池
}

type SfConf struct {
	StartTime string `yaml:"start_time" mapstructure:"start_time"` //时间序列开始时间
	MachineID int64  `yaml:"machine_id" mapstructure:"machine_id"` //机器ID
}

// 提供给全局访问配置的一个方法
func GetGlobalConf() *GlobalConfig {
	once.Do(readConfig)
	return &config
}

// 使用Viper管理配置
func readConfig() {
	viper.SetConfigName("config") //指定配置文件的名称（不需要带后缀）
	viper.SetConfigType("yml")    //指定配置文件类型
	viper.AddConfigPath("./conf") //指定查找配置文件的路径，这里使用相对路径
	viper.AddConfigPath(".")
	viper.AddConfigPath("../conf")
	err := viper.ReadInConfig() //读取配置信息
	if err != nil {
		//读取配置信息失败
		panic("viper.ReadInConfig() failed, err:" + err.Error())
	}
	//将读取到得文件反序列化到 config变量中
	if err = viper.Unmarshal(&config); err != nil {
		panic("viper.Unmarshal(&config) failed, err:" + err.Error())
	}
	//支持配置修改后自动加载配置信息
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件修改了...")
	})
}
