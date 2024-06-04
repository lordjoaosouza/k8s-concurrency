package logger

import (
	"github.com/lordjoaosouza/k3s-concurrency/config"
	"github.com/rs/zerolog"
	"os"
)

func New(cfg *config.Config) *zerolog.Logger {
	logger := zerolog.New(os.Stderr).With().Str("service", cfg.InternalConfig.ServiceName).Timestamp().Logger()

	return &logger
}
