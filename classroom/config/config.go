package config

import (
	"time"

	"github.com/spf13/viper"
)

const GRPC_PORT = "GRPC_PORT"

type Config struct {
	Server
	Telemetry
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
