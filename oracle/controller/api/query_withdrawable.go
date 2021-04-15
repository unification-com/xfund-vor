package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (d *Oracle) QueryWithdrawableTokens(c echo.Context) error {
	withdrawable, err := d.service.QueryWithdrawableTokens()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, withdrawable)
}
