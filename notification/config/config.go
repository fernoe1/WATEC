package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Nats
	Mailjet
}

type Nats struct {
	Addr                string
	NotificationSubject string
}

type Mailjet struct {
	ApiKey    string
	ApiSecret string
	FromEmail string
	FromName  string
}

func ReadConfig() (*Config, error) {
	viper.SetConfigName("config.local")
	viper.AddConfigPath("../config")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
