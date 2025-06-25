package limiter

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"

	"github.com/go-redis/redis_rate/v10"
)

type RateLimiter struct {
	redisClient *redis.Client
	limiter     *redis_rate.Limiter
}

type IRateLimit interface {
	AllowPerSec(key string, limit int) (*redis_rate.Result, error)
	LimitRequestPerSecond(key string, limit int) (bool, error)
}

var RateLimit IRateLimit
var ctx = context.Background()

func NewRateLimiter(host, pass string) IRateLimit {
	rdb := redis.NewClient(&redis.Options{
		Addr:         host,
		Password:     pass,
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

func (r *RateLimiter) AllowPerSec(key string, limit int) (*redis_rate.Result, error) {
	res, err := r.limiter.Allow(ctx, key, redis_rate.PerSecond(limit))
	return res, err
}

func (r *RateLimiter) LimitRequestPerSecond(key string, limit int) (bool, error) {
	if res, err := r.limiter.Allow(ctx, key, redis_rate.PerSecond(limit)); err != nil {
		return false, err
	} else if res.Allowed < 1 {
		return true, nil
	} else {
		return false, nil
	}
}
