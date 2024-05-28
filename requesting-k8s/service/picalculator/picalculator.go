package picalculator

import (
	"context"
	"github.com/lordjoaosouza/k3s-concurrency/presenter/res"
	"github.com/rs/zerolog"
	"os"
)

type Service struct {
	logger *zerolog.Logger
}

func New(logger *zerolog.Logger) *Service {
	return &Service{
		logger: logger,
	}
}

func (svc Service) CalculatePI(ctx context.Context, quantity int) (*res.CalculatePi, error) {
	hostName, err := os.Hostname()
	if err != nil {
		return nil, err
	}

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

	return &res.CalculatePi{
		HostName: hostName,
		PiValue:  pi,
	}, nil
}
