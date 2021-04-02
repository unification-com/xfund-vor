package api

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"net/http"
	"oracle/models/api"
)

func (d *Oracle) Register(c echo.Context) error {
	var requestModel api.OracleRegisterRequestModel
	json.NewDecoder(c.Request().Body).Decode(&requestModel)
	transactionInfo, err := d.service.Register(requestModel.AccountName, requestModel.PrivateKey, requestModel.Fee, requestModel.ProviderPaysGas)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, transactionInfo)
}
