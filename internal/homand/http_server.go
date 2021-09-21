package homand

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type server struct {
	hc HealthChecker
}

func NewServer(hc HealthChecker) *server {
	return &server{hc: hc}
}

func (s server) CheckHealth(c echo.Context) error {
	var req CheckHealthJSONRequestBody
	err := c.Bind(&req)
	if err != nil {
		return err
	}

	if string(req.CheckType) != string(HealthCheckHTTP) && string(req.CheckType) != string(HealthCheckTCP) {
		return c.JSON(http.StatusBadRequest, GeneralRes{
			Code:    http.StatusBadRequest,
			Message: "Unkown check_type",
		})
	}
	if req.Address == "" {
		return c.JSON(http.StatusBadRequest, GeneralRes{
			Code:    http.StatusBadRequest,
			Message: "Unknown address",
		})
	}

	isAvailable, err := s.hc.IsAvailable(c.Request().Context(), HealthCheckType(req.CheckType), req.Address)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, GeneralRes{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(200, CheckHealthRes{IsAvailable: isAvailable})
}
