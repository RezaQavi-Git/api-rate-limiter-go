package main

import (
	"api-rate-limiter-go/configs"
	"api-rate-limiter-go/internal"
	"net/http"
)

func main() {
	println("Rate Limiter Application Started...")
	appConfig := configs.Load()

	rateLimiter := internal.NewRateLimiter(appConfig.RateLimitConfigs)
	http.Handle("/api/", rateLimiter.HttpMiddleware(http.HandlerFunc(handleAPI)))
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
