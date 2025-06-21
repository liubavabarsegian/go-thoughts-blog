package rate_limiter

import (
	"context"
	"time"

	"github.com/go-redis/redis_rate/v10"
	"github.com/redis/go-redis/v9"
)

type RateLimiter struct {
	client  *redis.Client
	limiter *redis_rate.Limiter
}

func NewRateLimiter() *RateLimiter {
	client := redis.NewClient(&redis.Options{
		Addr: "redis:6379",
	})
	limiter := redis_rate.NewLimiter(client)

	return &RateLimiter{
		client:  client,
		limiter: limiter,
	}
}

func (r *RateLimiter) Limit(ctx context.Context, key string, limit int, period time.Duration) (*redis_rate.Result, error) {
	return r.limiter.Allow(ctx, key, redis_rate.Limit{
		Rate:   limit,  // сколько разрешается
		Period: period, // за какой период
		Burst:  limit,  // максимальное "окно"
	})
}
