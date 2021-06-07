package ratelimiter

import "time"

type RateLimiter interface {
	WaitMaxDuration(cnt int, maxWait time.Duration) bool
	Rate() float64
}
