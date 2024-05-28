package picalculator

import (
	"github.com/lordjoaosouza/k3s-concurrency/util/resutil"
	"github.com/rs/zerolog"
)

type Service struct {
	resutil *resutil.ResUtil
	logger  *zerolog.Logger
}

func New(logger *zerolog.Logger, util *resutil.ResUtil) *Service {
	return &Service{
		resutil: util,
		logger:  logger,
	}
}
