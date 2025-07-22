
package utils

import (
    "time"

    "github.com/patrickmn/go-cache"
)

type RateLimiter struct {
    limiter *cache.Cache
    ttl     time.Duration
}

func NewRateLimiter(ttl time.Duration) *RateLimiter {
    return &RateLimiter{
        limiter: cache.New(ttl, time.Minute),
        ttl:     ttl,
    }
}

func (r *RateLimiter) Allow(phone string) bool {
    _, found := r.limiter.Get(phone)
    if found {
        return false // هنوز در محدودیت زمانی هست
    }
    r.limiter.Set(phone, true, r.ttl)
    return true
}
