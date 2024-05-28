package config

import (
	"github.com/spf13/viper"
)

type InternalConfig struct {
	RunningLocal bool
	ServerPort   int
	ServiceName  string
}

type Config struct {
	InternalConfig *InternalConfig
}

func Get() *Config {
	viper.AutomaticEnv()

	return &Config{
		InternalConfig: &InternalConfig{
			RunningLocal: true,
			ServerPort:   8080,
			ServiceName:  "requesting-k8s",
		},
	}
}
