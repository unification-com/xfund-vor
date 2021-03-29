package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"oracle/models/api"
)

func (d *Oracle) Register(c echo.Context) error {
	var requestModel api.OracleRegisterRequestModel
	c.Bind(&requestModel)
	transactionInfo, err := d.service.Register(requestModel.AccountName, requestModel.PrivateKey, requestModel.Fee, requestModel.ProviderPaysGas)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, transactionInfo)
}
