package service

import (
	"context"
	"github.com/lordjoaosouza/k3s-concurrency/config"
	"github.com/lordjoaosouza/k3s-concurrency/service/picalculator"
	"github.com/lordjoaosouza/k3s-concurrency/util/resutil"
	"github.com/rs/zerolog"
)

type Service struct {
	PiCalculatorService *picalculator.Service
}

func New(cfg *config.Config, logger *zerolog.Logger, util *resutil.ResUtil) *Service {

	return &Service{
		PiCalculatorService: picalculator.New(logger, util),
	}
}

func (svc *Service) CalculatePI(ctx context.Context, quantity int) (float64, error) {
	pi := 0.0
	denominator := 1.0
	for i := 0; i < quantity; i++ {
		if i%2 == 0 {
			pi += 4 / denominator
		} else {
			pi -= 4 / denominator
		}

		denominator += 2.0
	}

	return pi, nil
}
