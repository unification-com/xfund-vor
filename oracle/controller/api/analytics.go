package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func (d *Oracle) Analytics(c echo.Context) error {
	xFundEth, _  := strconv.ParseFloat(c.QueryParam("eth"), 64)
	xFundUsd, _  := strconv.ParseFloat(c.QueryParam("usd"), 64)
	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	gasPrice, _ := strconv.Atoi(c.QueryParam("gasprice"))
	fees, _  := strconv.ParseFloat(c.QueryParam("fees"), 64)
	simulation, _ := strconv.Atoi(c.QueryParam("sim"))
	consumer := c.QueryParam("consumer")

	analytics, err := d.service.Analytics(xFundEth, xFundUsd, fees, limit, int64(gasPrice), simulation, consumer)

	if err != nil {
		return c.JSONPretty(http.StatusInternalServerError, analytics, "  ")
	}
	return c.JSONPretty(http.StatusOK, analytics, "  ")
}
