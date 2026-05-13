package redis

import (
	"context"
	"log"
	"time"

	"github.com/fernoe1/WATEC/gateway/config"
	"github.com/redis/go-redis/extra/redisotel/v9"
	"github.com/redis/go-redis/v9"
)

func NewRedis(cfg *config.Config) *redis.Client {
	r := redis.NewClient(&redis.Options{
		Addr:       cfg.Redis.Addr,
		Password:   cfg.Redis.Password,
		DB:         cfg.Redis.DB,
		MaxRetries: cfg.Redis.MaxRetries,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 7*time.Second)
	defer cancel()

	if err := r.Ping(ctx).Err(); err != nil {
		log.Fatal(err)
	}

	if err := redisotel.InstrumentTracing(r); err != nil {
		log.Fatal(err)
	}

	if err := redisotel.InstrumentMetrics(r); err != nil {
		log.Fatal(err)
	}

	return r
}
