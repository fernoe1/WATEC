package config

import (
	"time"

	"github.com/spf13/viper"
)

const GRPC_PORT = "GRPC_PORT"

type Config struct {
	Server
	Telemetry
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

type Telemetry struct {
	Name     string
	Endpoint string
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

	var c Config
	if err := viper.Unmarshal(&c); err != nil {
		return nil, err
	}

	return &c, nil
}
