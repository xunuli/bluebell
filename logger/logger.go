package logger

import (
	"bluebell/config"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
	"time"
)

var log *zap.Logger

// 自定义zap日志需要实现 Encoder、WriterSyncer、LogLevel
// InitLog 初始化日志
func InitLog(cfg *config.LogConf, mode string) {
	//指定日志写到那里去，调用zap.WriteSyncer()方法
	writeSyncer := getLogWriter(cfg.Filename, cfg.LogPath, cfg.MaxSize, cfg.MaxBackups, cfg.MaxAge)
	encoder := getEncoder()
	var le = new(zapcore.Level)
	err := le.UnmarshalText([]byte(cfg.Lever))
	if err != nil {
		fmt.Printf("le.UnmarshalText([]byte(cfg.Lever)) failed, err:%s", err)
		return
	}
	var core zapcore.Core
	if mode == "debug" {
		//进入开发模式，日志输出到终端和日志文件
		consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
		core = zapcore.NewTee(
			//输出日志到文件中
			zapcore.NewCore(encoder, writeSyncer, le),
			//输出日志到终端
			zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stdout), zapcore.DebugLevel),
		)
	} else {
		//输出日志到文件中
		zapcore.NewCore(encoder, writeSyncer, le)
	}
	//构造一个log
	log = zap.New(core, zap.AddCaller())
	//将log替换为全局的log
	zap.ReplaceGlobals(log)
	zap.L().Info("init logger success")
}

// 如何写入日志
func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

// 指定日志写到哪里去
func getLogWriter(filename, logpath string, maxSize, maxBackup, maxAge int) zapcore.WriteSyncer {
	logPath := logpath
	//lumberjack用于按切割日志文件
	lumberjackLogger := &lumberjack.Logger{
		Filename:   logPath + "/" + filename,
		MaxAge:     maxAge,
		MaxBackups: maxBackup,
		MaxSize:    maxSize,
	}
	//AddSync将打开的文件句柄传进去
	return zapcore.AddSync(lumberjackLogger)
}

// 以中间件的形式将zap配置为Gin框架的默认日志
func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path      //请求路径
		query := c.Request.URL.RawQuery //请求
		//向后继续执行其他中间件或路由处理函数
		c.Next()

		cost := time.Since(start)
		//打印日志信息
		log.Info(path,
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.Duration("cost", cost),
		)
	}
}

// GinRecvery recover掉项目中可能出现的panic，并使用zap日志记录
func GinRecovery(stack bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			//捕获异常panic
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					log.Error(c.Request.URL.Path,
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
					// If the connection is dead, we can't write a status to it.
					c.Error(err.(error)) // nolint: errcheck
					c.Abort()
					return
				}

				if stack {
					log.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
						zap.String("stack", string(debug.Stack())),
					)
				} else {
					log.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
				}
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}
