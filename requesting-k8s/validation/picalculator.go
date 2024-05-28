package validation

import (
	"errors"
	"github.com/labstack/echo/v4"
	"strconv"
)

func ValidateCalculatePi(ctx echo.Context) (int, error) {
	quantityStr := ctx.Param("quantityInt")

	if quantityStr == "" {
		return 0, errors.New("quantityInt variable is null")
	}

	quantity, err := strconv.Atoi(quantityStr)
	if err != nil {
		return 0, errors.New(err.Error())
	}

	return quantity, nil
}
