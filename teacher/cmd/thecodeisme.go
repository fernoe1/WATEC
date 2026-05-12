package main

import (
	"log"
	"log/slog"

	"github.com/fernoe1/WATEC/teacher/config"
	"github.com/fernoe1/WATEC/teacher/internal/server"
	"github.com/fernoe1/WATEC/teacher/migrate"
	"github.com/fernoe1/WATEC/teacher/pkg/gorm"
	"github.com/fernoe1/WATEC/teacher/pkg/redis"
)

func main() {
	slog.Info("reading config")
	cfg, err := config.ReadConfig()
	if err != nil {
		log.Fatal(err)
	}

	gormDB := gorm.NewGormDB(cfg)
	slog.Info("gorm connected")
	migrate.Migrate(gormDB)

	redisClient := redis.NewRedis(cfg)
	slog.Info("redis connected")

	s := server.NewServer(gormDB, redisClient, cfg)
	log.Fatal(s.Run())
}
