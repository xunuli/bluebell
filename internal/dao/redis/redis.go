package redis

import (
	"bluebell/config"
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

var (
	redisClient *redis.Client
)

func InitRedis(cfg *config.RedisConf) {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password: cfg.Password,
		DB:       cfg.Db,
		PoolSize: cfg.Poolsize,
	})
	if redisClient == nil {
		panic("failed to call redis.NewClient")
	}

	pong, err := redisClient.Ping(context.Background()).Result()
	fmt.Println(pong)
	if err != nil {
		panic("Failed to ping redis, err:%s" + err.Error())
	}
}
