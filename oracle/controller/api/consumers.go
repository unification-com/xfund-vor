package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func (d *Oracle) Consumers(c echo.Context) error {
	xFundEth, _ := strconv.ParseFloat(c.QueryParam("eth"), 64)
	xFundUsd, _ := strconv.ParseFloat(c.QueryParam("usd"), 64)
	consumer := c.QueryParam("consumer")
	analyticsData, err := d.service.Consumers(xFundEth, xFundUsd, consumer)

	if err != nil {
		return c.JSONPretty(http.StatusInternalServerError, analyticsData, "  ")
	}
	return c.JSONPretty(http.StatusOK, analyticsData, "  ")
}
