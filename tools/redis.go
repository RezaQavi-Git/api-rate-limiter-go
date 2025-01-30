package tools

import (
	"api-rate-limiter-go/configs"
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

func NewRedisClient(config *configs.RedisConfigs) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:            config.Address,
		DialTimeout:     time.Duration(config.DialTimeout) * time.Second,
		ReadTimeout:     time.Duration(config.ReadTimeout) * time.Second,
		WriteTimeout:    time.Duration(config.WriteTimeout) * time.Second,
		ConnMaxIdleTime: time.Duration(config.MaxIdleTime) * time.Second,
		PoolSize:        config.PoolSize,
	})

	pong, err := rdb.Ping(context.Background()).Result()
	fmt.Println(pong, err)
	return rdb
}
