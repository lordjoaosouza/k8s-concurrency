package controller

import (
	"github.com/lordjoaosouza/k3s-concurrency/server/controller/picalculator"
	"github.com/lordjoaosouza/k3s-concurrency/service"
	"github.com/lordjoaosouza/k3s-concurrency/util/resutil"
	"github.com/rs/zerolog"
)

type Controller struct {
	PiCalculatorController *picalculator.Controller
}

func New(svc *service.Service, logger *zerolog.Logger) *Controller {
	util := resutil.New(logger)

	return &Controller{
		PiCalculatorController: picalculator.New(svc, logger, util),
	}
}
