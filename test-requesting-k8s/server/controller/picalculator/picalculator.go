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

	PiCalculatorResponse, err := ctrl.svc.PiCalculatorService.CalculatePI(ctx.Request().Context(), quantity)
	if err != nil {
		return ctx.JSON(ctrl.resutil.Wrap(nil, err, http.StatusInternalServerError))
	}

	return ctx.JSON(ctrl.resutil.Wrap(PiCalculatorResponse, nil, http.StatusOK))
}

func (ctrl *Controller) TestAPI(ctx echo.Context) error {
	testStruct, err := validation.ValidateTestAPI(ctx.Request().Body)
	if err != nil {
		return ctx.JSON(ctrl.resutil.Wrap(nil, err, http.StatusBadRequest))
	}

	time, err := ctrl.svc.PiCalculatorService.TestAPI(ctx.Request().Context(), testStruct)
	if err != nil {
		return ctx.JSON(ctrl.resutil.Wrap(nil, err, http.StatusBadRequest))
	}

	return ctx.JSON(ctrl.resutil.Wrap(time, nil, http.StatusOK))
}
