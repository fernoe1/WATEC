package repository

import (
	"context"
	"encoding/json"
	"log/slog"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisRepository struct {
	log *slog.Logger
	r   *redis.Client
}

func NewRedisRepository(log *slog.Logger, r *redis.Client) *RedisRepository {
	return &RedisRepository{log: log, r: r}
}

func (r *RedisRepository) getTTLUntilNextHour() time.Duration {
	now := time.Now().In(time.FixedZone("UTC+6", 6*3600))
	next := now.Add(time.Hour).Truncate(time.Hour)

	return next.Sub(now)
}

func (r *RedisRepository) Set(ctx context.Context, free []int64) error {
	if free == nil {
		free = []int64{}
	}

	data, err := json.Marshal(free)
	if err != nil {

		return err
	}

	return r.r.Set(ctx, "free", data, r.getTTLUntilNextHour()).Err()
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

func (r *RedisRepository) Del(ctx context.Context) error {
	return r.r.Del(ctx, "free").Err()
}
