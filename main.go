package main

import (
	"api-rate-limiter-go/internal"
	"net/http"
	"time"
)

func main() {
	println("Rate Limiter Application Started...")
	configs := []internal.RateLimitConfig{
		{
			Pattern:  `^/api/.*`,
			Limit:    5,
			Duration: 10 * time.Second,
		},
	}

	rateLimiter := internal.NewRateLimiter(configs)
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
