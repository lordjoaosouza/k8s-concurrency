package router

import (
	"github.com/labstack/echo/v4"
	"github.com/lordjoaosouza/k3s-concurrency/config"
	"github.com/lordjoaosouza/k3s-concurrency/server/controller"
	"github.com/lordjoaosouza/k3s-concurrency/server/controller/healthcheck"
)

func Register(cfg *config.Config, svr *echo.Echo, ctrl *controller.Controller) {
	root := svr.Group("")
	root.GET("/test", healthcheck.HealthCheck)

	piCalculator := root.Group("/picalc")
	piCalculator.GET("/:quantity", ctrl.PiCalculatorController.CalculatePi)
	piCalculator.POST("/testapi", ctrl.PiCalculatorController.TestAPI)

}
