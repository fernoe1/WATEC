package repository

import (
	"context"

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
	//TODO implement me
	panic("implement me")
}

func (r *RedisRepository) Get(ctx context.Context, number int64) (*domain.Locker, error) {
	//TODO implement me
	panic("implement me")
}
