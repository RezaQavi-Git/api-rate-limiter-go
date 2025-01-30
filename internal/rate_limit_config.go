package internal

import (
	"regexp"
	"time"
)

type RateLimitConfig struct {
	Pattern  string         `json:"pattern"`
	Regex    *regexp.Regexp `json:"regex"`
	Limit    int            `json:"limit"`
	Duration time.Duration  `json:"duration"`
}
