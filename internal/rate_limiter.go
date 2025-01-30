package internal

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"net/http"
	"regexp"
	"time"
)

type RateLimiter struct {
	configs     []RateLimitConfig
	redisClient *redis.Client
}

func NewRateLimiter(configs []RateLimitConfig, client *redis.Client) *RateLimiter {
	for i := range configs {
		configs[i].Regex = regexp.MustCompile(configs[i].Pattern)
	}
	return &RateLimiter{
		configs:     configs,
		redisClient: client,
	}
}

func (r *RateLimiter) HttpMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		basePath := req.URL.Path
		userKey := req.Host

		config, matched := r.matchConfig(basePath)
		if !matched {
			next.ServeHTTP(w, req)
			return
		}

		if !r.allowRequest(userKey, config) {
			http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, req)
	})
}

func (r *RateLimiter) allowRequest(userKey string, config *RateLimitConfig) bool {
	now := time.Now().Unix()
	window := now / int64(time.Duration(config.Duration)*time.Second)

	redisKey := fmt.Sprintf("%s:%d", userKey, window)

	count, err := r.redisClient.Incr(context.Background(), redisKey).Result()
	if err != nil {
		return false
	}

	if count == 1 {
		r.redisClient.Expire(context.Background(), redisKey, time.Duration(config.Duration)*time.Second)
	}

	println("cnt: ", count)
	return count <= int64(config.Limit)
}

func (r *RateLimiter) matchConfig(path string) (*RateLimitConfig, bool) {
	for _, config := range r.configs {
		matched := config.Regex.MatchString(path)
		if matched {
			return &config, true
		}
	}
	return nil, false
}
