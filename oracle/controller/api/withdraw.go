package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"oracle/models/api"
)

func (d *Oracle) Withdraw(c echo.Context) error {
	var requestModel api.OracleWithdrawRequestModel
	c.Bind(&requestModel)
	transactionInfo, err := d.service.Oracle.Withdraw(requestModel.Address, requestModel.Amount)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, transactionInfo)
	}
	return c.JSON(http.StatusOK, transactionInfo)
}
