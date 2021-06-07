package ratelimiter

import (
	"github.com/juju/ratelimit"
	"time"
)

type Limiter struct {
	bucket *ratelimit.Bucket
}

func (l *Limiter) WaitMaxDuration(cnt int, maxWait time.Duration) bool {
	return l.bucket.WaitMaxDuration(int64(cnt), maxWait)
}

func (l *Limiter) Rate() float64 {
	return l.bucket.Rate()
}

func NewLimiter(qps int64) RateLimiter {
	interval := time.Second / time.Duration(qps)
	capacity := int64(float64(qps) * 1.2)

	return &Limiter{
		bucket: ratelimit.NewBucket(interval, capacity),
	}
}

//根据每分钟的请求量创建限流器
func NewLimiterByQpm(qpm int64) RateLimiter {
	interval := time.Minute / time.Duration(qpm)
	capacity := int64(float64(qpm) * 1.2)

	return &Limiter{
		bucket: ratelimit.NewBucket(interval, capacity),
	}
}
