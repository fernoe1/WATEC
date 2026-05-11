package redis

import (
	"log"

	"github.com/fernoe1/WATEC/classroom/config"
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

	if err := redisotel.InstrumentTracing(r); err != nil {
		log.Fatal(err)
	}

	if err := redisotel.InstrumentMetrics(r); err != nil {
		log.Fatal(err)
	}

	return r
}
