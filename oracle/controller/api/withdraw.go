package api

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"

	"oracle/models/api"
)

func (d *Oracle) Withdraw(c echo.Context) error {
	var requestModel api.OracleWithdrawRequestModel
	if err := decodeBody(c, &requestModel); err != nil {
		return badRequest(c, err)
	}
	_, addressErr := parseAddress("address", requestModel.Address)
	if err := errors.Join(addressErr, positive("amount", requestModel.Amount)); err != nil {
		return badRequest(c, err)
	}
	transactionInfo, err := d.service.Withdraw(requestModel.Address, requestModel.Amount)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, transactionInfo)
}
