package api

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"net/http"
	"oracle/models/api"
)

func (d *Oracle) Withdraw(c echo.Context) error {
	var requestModel api.OracleWithdrawRequestModel
	json.NewDecoder(c.Request().Body).Decode(&requestModel)
	transactionInfo, err := d.service.Withdraw(requestModel.Address, requestModel.Amount)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, transactionInfo)
}
