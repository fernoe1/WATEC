package config

import (
	"strings"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Http
	Telemetry
}

type Http struct {
	Port              string
	PprofPort         string
	ReadHeaderTimeout time.Duration
	ReadTimeout       time.Duration
	WriteTimeout      time.Duration
	IdleTimeout       time.Duration
	MaxHeaderBytes    int
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

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	if !strings.HasPrefix(cfg.Http.Port, ":") {
		cfg.Http.Port = ":" + cfg.Http.Port
	}

	if !strings.HasPrefix(cfg.Http.PprofPort, ":") {
		cfg.Http.PprofPort = ":" + cfg.Http.PprofPort
	}

	return &cfg, nil
}
