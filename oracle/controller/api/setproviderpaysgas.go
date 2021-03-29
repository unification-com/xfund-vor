package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"oracle/models/api"
)

func (d *Oracle) SetProviderPaysGas(c echo.Context) error {
	var requestModel api.OracleSetProviderPaysGasRequestModel
	c.Bind(&requestModel)
	transactionInfo, err := d.service.SetProviderPaysGas(requestModel.ProviderPays)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, transactionInfo)
	}
	return c.JSON(http.StatusOK, transactionInfo)
}
