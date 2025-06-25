package limiter

import (
	"context"
	"time"

	"github.com/go-redis/redis_rate/v10"
	"github.com/redis/go-redis/v9"
)

type RateLimiter struct {
	redisClient *redis.Client
	limiter     *redis_rate.Limiter
}

type IRateLimiter interface {
	AllowPerSecord(key string, second int) (*redis_rate.Result, error)
	LimitRequestPerSecord(key string, limit int) (bool, error)
}

var RateLimit IRateLimiter
var ctx = context.Background()

func NewRateLimiter(host string, auth string) IRateLimiter {
	rdb := redis.NewClient(&redis.Options{
		Addr:         host,
		Password:     auth,
		DB:           10,
		PoolSize:     30,
		PoolTimeout:  20 * time.Second,
		ReadTimeout:  20 * time.Second,
		WriteTimeout: 15 * time.Second,
	})
	rateLimit := &RateLimiter{redisClient: rdb}
	rateLimit.SetupLimiter()
	return rateLimit
}

func (rateLimit *RateLimiter) SetupLimiter() {
	_ = rateLimit.redisClient.FlushDB(ctx).Err()
	limiter := redis_rate.NewLimiter(rateLimit.redisClient)
	rateLimit.limiter = limiter
}

func (rateLimit *RateLimiter) AllowPerSecord(key string, limit int) (*redis_rate.Result, error) {
	res, err := rateLimit.limiter.Allow(ctx, key, redis_rate.PerSecond(limit))
	return res, err
}
func (rateLimit *RateLimiter) LimitRequestPerSecord(key string, limit int) (bool, error) {
	if res, err := rateLimit.limiter.Allow(ctx, key, redis_rate.PerSecond(limit)); err != nil {
		return false, err
	} else if res.Allowed < 1 {
		return true, nil
	} else {
		return false, nil
	}
}
