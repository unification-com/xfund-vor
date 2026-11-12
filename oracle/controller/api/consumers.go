package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"oracle/service"
	"strconv"
)

func (d *Oracle) Consumers(c echo.Context, svc *service.Service) error {
	xFundEth, _ := strconv.ParseFloat(c.QueryParam("eth"), 64)
	xFundUsd, _ := strconv.ParseFloat(c.QueryParam("usd"), 64)
	consumer := c.QueryParam("consumer")
	analyticsData, err := svc.Consumers(xFundEth, xFundUsd, consumer)

	if err != nil {
		return c.JSONPretty(http.StatusInternalServerError, analyticsData, "  ")
	}
	return c.JSONPretty(http.StatusOK, analyticsData, "  ")
}
