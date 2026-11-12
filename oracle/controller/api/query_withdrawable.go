package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"oracle/service"
)

func (d *Oracle) QueryWithdrawableTokens(c echo.Context, svc *service.Service) error {
	withdrawable, err := svc.QueryWithdrawableTokens()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, withdrawable)
}
