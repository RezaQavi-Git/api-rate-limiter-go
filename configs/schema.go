package configs

import (
	"api-rate-limiter-go/internal"
)

type Config struct {
	RedisConfigs     *RedisConfigs              `mapstructure:"redis_configs"`
	RateLimitConfigs []internal.RateLimitConfig `mapstructure:"rate_limit_configs"`
}

type RedisConfigs struct {
	Addresses                 []string `mapstructure:"addresses"`
	MasterName                string   `mapstructure:"master_name"`
	ConnectionPoolMaxIdleTime int      `mapstructure:"connection_pool_max_idle_time"`
	ReadTimeout               int      `mapstructure:"read_timeout"`
}
