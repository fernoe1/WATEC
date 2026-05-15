package repository

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/fernoe1/WATEC/locker/internal/domain"
	"github.com/redis/go-redis/v9"
)

type RedisRepository struct {
	r *redis.Client
}

func NewRedisRepository(r *redis.Client) *RedisRepository {
	return &RedisRepository{r: r}
}

func (r *RedisRepository) Set(ctx context.Context, locker *domain.Locker) error {
	data, err := json.Marshal(locker)
	if err != nil {
		return err
	}

	key := fmt.Sprintf("locker:%d", locker.Number)

	return r.r.Set(ctx, key, data, 0).Err()
}

func (r *RedisRepository) Get(ctx context.Context, number int64) (*domain.Locker, error) {
	key := fmt.Sprintf("locker:%d", number)

	data, err := r.r.Get(ctx, key).Bytes()
	if err != nil {
		return nil, err
	}

	var locker domain.Locker

	err = json.Unmarshal(data, &locker)
	if err != nil {
		return nil, err
	}

	return &locker, nil
}

func (r *RedisRepository) Delete(ctx context.Context, number int64) error {
	key := fmt.Sprintf("locker:%d", number)

	return r.r.Del(ctx, key).Err()
}
