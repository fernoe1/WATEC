package repository

import (
	"context"

	"github.com/fernoe1/WATEC/teacher/internal/domain"
	"github.com/redis/go-redis/v9"
)

type RedisRepository struct {
	r *redis.Client
}

func NewRedisRepository(r *redis.Client) *RedisRepository {
	return &RedisRepository{r: r}
}

func (r RedisRepository) Get(ctx context.Context, name string) (error, *domain.Free) {
	//TODO implement me
	panic("implement me")
}

func (r RedisRepository) Set(ctx context.Context, teacher *domain.Teacher) error {
	//TODO implement me
	panic("implement me")
}
