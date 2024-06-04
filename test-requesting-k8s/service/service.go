package service

import (
	"github.com/lordjoaosouza/k3s-concurrency/config"
	"github.com/lordjoaosouza/k3s-concurrency/service/picalculator"
	"github.com/rs/zerolog"
)

type Service struct {
	PiCalculatorService *picalculator.Service
}

func New(cfg *config.Config, logger *zerolog.Logger) *Service {

	return &Service{
		PiCalculatorService: picalculator.New(logger),
	}
}
