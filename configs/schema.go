package configs

import (
	"api-rate-limiter-go/internal"
)

type Config struct {
	RedisConfigs     *RedisConfigs              `mapstructure:"redis_configs"`
	RateLimitConfigs []internal.RateLimitConfig `mapstructure:"rate_limit_configs"`
}

type RedisConfigs struct {
	Address      string `mapstructure:"address"`
	DialTimeout  int    `mapstructure:"dial_timeout"`
	ReadTimeout  int    `mapstructure:"read_timeout"`
	WriteTimeout int    `mapstructure:"write_timeout"`
	PoolSize     int    `mapstructure:"pool_size"`
	MaxIdleTime  int    `mapstructure:"max_idle_time"`
}
