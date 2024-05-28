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
			RunningLocal: viper.GetBool("RUNNING_LOCAL"),
			ServerPort:   viper.GetInt("SERVER_PORT"),
			ServiceName:  viper.GetString("SERVICE_NAME"),
		},
	}
}
