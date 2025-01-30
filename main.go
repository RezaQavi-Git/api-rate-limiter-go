package main

import (
	"api-rate-limiter-go/configs"
	"api-rate-limiter-go/internal"
	"api-rate-limiter-go/tools"
	"net/http"
)

func main() {
	println("Rate Limiter Application Started...")
	appConfig := configs.Load()
	rc := tools.NewRedisClient(appConfig.RedisConfigs)

	rateLimiter := internal.NewRateLimiter(appConfig.RateLimitConfigs, rc)
	http.Handle("/api/v2/", rateLimiter.HttpMiddleware(http.HandlerFunc(handleAPI)))
	http.Handle("/api/v3/", rateLimiter.HttpMiddleware(http.HandlerFunc(handleAPI)))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}

func handleAPI(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("OK"))
	if err != nil {
		return
	}
}
