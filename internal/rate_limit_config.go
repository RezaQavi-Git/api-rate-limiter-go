package internal

import (
	"regexp"
)

type RateLimitConfig struct {
	Pattern  string         `json:"pattern"`
	Regex    *regexp.Regexp `json:"regex"`
	Limit    int            `json:"limit"`
	Duration int            `json:"duration"`
}
