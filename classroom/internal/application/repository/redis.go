package repository

import (
	"context"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisRepository struct {
	r *redis.Client
}

func (r *RedisRepository) getTTLUntilNextHour() time.Duration {
	now := time.Now().In(time.FixedZone("UTC+6", 6*3600))
	next := now.Add(time.Hour).Truncate(time.Hour)

	return next.Sub(now)
}

func (r *RedisRepository) Set(ctx context.Context, free []int64) {
	r.r.Set(ctx, "free", free, r.getTTLUntilNextHour())
}

func (r *RedisRepository) Get(ctx context.Context) ([]int64, error) {
	result, err := r.r.Get(ctx, "free").Bytes()

	if err == redis.Nil {

		return nil, nil
	}

	if err != nil {

		return nil, err
	}

	var res []int64
	if err := json.Unmarshal(result, &res); err != nil {
		return nil, err
	}

	return res, nil
}
