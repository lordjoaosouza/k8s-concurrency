package validation

import (
	"encoding/json"
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/lordjoaosouza/k3s-concurrency/presenter/req"
	"io"
	"strconv"
)

func ValidateCalculatePi(ctx echo.Context) (int, error) {
	quantityStr := ctx.Param("quantity")

	if quantityStr == "" {
		return 0, errors.New("quantityInt variable is null")
	}

	quantity, err := strconv.Atoi(quantityStr)
	if err != nil {
		return 0, errors.New(err.Error())
	}

	return quantity, nil
}

func ValidateTestAPI(rc io.ReadCloser) (r *req.TestAPIRequest, err error) {
	defer rc.Close()

	body, err := io.ReadAll(rc)
	if err != nil {
		return nil, errors.New("cannot read requisition body")
	}

	if err := json.Unmarshal(body, &r); err != nil {
		return nil, err
	}

	return r, nil
}
