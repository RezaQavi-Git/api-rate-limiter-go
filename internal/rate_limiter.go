package internal

import (
	"log"
	"net/http"
	"regexp"
)

type RateLimiter struct {
	configs []RateLimitConfig
}

func NewRateLimiter(configs []RateLimitConfig) *RateLimiter {
	for i := range configs {
		configs[i].Regex = regexp.MustCompile(configs[i].Pattern)
	}
	return &RateLimiter{configs: configs}
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
	log.Println("userKey:", userKey, "config:", *config, "was accepted")
	return true
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
