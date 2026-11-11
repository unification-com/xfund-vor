package api

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"

	"oracle/models/api"
	"oracle/store/keystorage"
)

func (d *Oracle) Register(c echo.Context) error {
	var requestModel api.OracleRegisterRequestModel
	if err := decodeBody(c, &requestModel); err != nil {
		return badRequest(c, err)
	}
	_, keyErr := keystorage.ParsePrivateKey(requestModel.PrivateKey)
	if err := errors.Join(required("account_name", requestModel.AccountName), keyErr, positive("fee", requestModel.Fee)); err != nil {
		return badRequest(c, err)
	}
	transactionInfo, err := d.service.Register(requestModel.AccountName, requestModel.PrivateKey, requestModel.Fee)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, transactionInfo)
}
