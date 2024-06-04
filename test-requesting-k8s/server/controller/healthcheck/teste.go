package healthcheck

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func HealthCheck(ctx echo.Context) error {

	return ctx.String(http.StatusOK, "ok")
}
