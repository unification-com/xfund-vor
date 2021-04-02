package api

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"net/http"
	"oracle/models/api"
)

func (d *Oracle) SetProviderPaysGas(c echo.Context) error {
	var requestModel api.OracleSetProviderPaysGasRequestModel
	json.NewDecoder(c.Request().Body).Decode(&requestModel)
	transactionInfo, err := d.service.SetProviderPaysGas(requestModel.ProviderPays)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, transactionInfo)
}
