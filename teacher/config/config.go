package config

import (
	"strings"
	"time"

	"github.com/spf13/viper"
)

const GRPC_PORT = "GRPC_PORT"

type Config struct {
	Server
	PostgreSQL
	Redis
}

type Server struct {
	Port              string
	Development       bool
	Timeout           time.Duration
	MaxConnectionIdle time.Duration
	MaxConnectionAge  time.Duration
}

type PostgreSQL struct {
	Host     string
	User     string
	Password string
	Port     string
	SslMode  string
	Timezone string
}

type Redis struct {
	Addr       string
	Password   string
	DB         int
	MaxRetries int
}

func ReadConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.AddConfigPath("../config")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	if !strings.HasPrefix(cfg.Server.Port, ":") {
		cfg.Server.Port = ":" + cfg.Server.Port
	}

	return &cfg, nil
}
