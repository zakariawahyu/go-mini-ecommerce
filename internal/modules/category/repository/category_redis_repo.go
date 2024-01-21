package repository

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

type categoryRedisRepo struct {
	redis *redis.Client
}

func NewCategoryRedisRepo(redis *redis.Client) *categoryRedisRepo {
	return &categoryRedisRepo{
		redis: redis,
	}
}

func (r *categoryRedisRepo) Set(ctx context.Context, key string, value interface{}, exp time.Duration) error {
	return r.redis.Set(ctx, key, value, exp).Err()
}

func (r *categoryRedisRepo) Get(ctx context.Context, key string) (string, error) {
	return r.redis.Get(ctx, key).Result()
}

func (r *categoryRedisRepo) Delete(ctx context.Context, key string) error {
	if err := r.redis.Del(ctx, key).Err(); err != nil {
		return err
	}
	return nil
}
