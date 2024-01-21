package repository

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

type productRedisRepo struct {
	redis *redis.Client
}

func NewProductRedisRepo(redis *redis.Client) *productRedisRepo {
	return &productRedisRepo{
		redis: redis,
	}
}

func (r *productRedisRepo) Set(ctx context.Context, key string, value interface{}, exp time.Duration) error {
	return r.redis.Set(ctx, key, value, exp).Err()
}

func (r *productRedisRepo) Get(ctx context.Context, key string) (string, error) {
	return r.redis.Get(ctx, key).Result()
}

func (r *productRedisRepo) Delete(ctx context.Context, key string) error {
	if err := r.redis.Del(ctx, key).Err(); err != nil {
		return err
	}
	return nil
}
