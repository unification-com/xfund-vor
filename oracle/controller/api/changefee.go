package api

import (
	"net/http"
	"oracle/service"

	"github.com/labstack/echo/v4"

	"oracle/models/api"
)

func (d *Oracle) ChangeFee(c echo.Context, svc *service.Service) error {
	var requestModel api.OracleChangeFeeRequestModel
	if err := decodeBody(c, &requestModel); err != nil {
		return badRequest(c, err)
	}
	if err := positive("amount", requestModel.Amount); err != nil {
		return badRequest(c, err)
	}
	transactionInfo, err := svc.ChangeFee(requestModel.Amount)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, transactionInfo)
}
