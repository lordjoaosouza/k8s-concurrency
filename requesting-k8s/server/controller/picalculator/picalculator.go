package picalculator

import (
	"github.com/labstack/echo/v4"
	"github.com/lordjoaosouza/k3s-concurrency/service"
	"github.com/lordjoaosouza/k3s-concurrency/util/resutil"
	"github.com/lordjoaosouza/k3s-concurrency/validation"
	"github.com/rs/zerolog"
	"net/http"
)

type Controller struct {
	svc     *service.Service
	logger  *zerolog.Logger
	resutil *resutil.ResUtil
}

func New(svc *service.Service, logger *zerolog.Logger, util *resutil.ResUtil) *Controller {
	return &Controller{
		svc:     svc,
		logger:  logger,
		resutil: util,
	}
}

func (ctrl *Controller) CalculatePi(ctx echo.Context) error {
	quantity, err := validation.ValidateCalculatePi(ctx)
	if err != nil {
		return ctx.JSON(ctrl.resutil.Wrap(nil, err, http.StatusBadRequest))
	}

	return
}
