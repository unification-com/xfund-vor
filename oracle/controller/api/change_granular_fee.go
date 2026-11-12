package api

import (
	"errors"
	"net/http"
	"oracle/service"

	"github.com/labstack/echo/v4"

	"oracle/models/api"
)

func (d *Oracle) ChangeGranularFee(c echo.Context, svc *service.Service) error {
	var requestModel api.OracleChangeGranularFeeRequestModel
	if err := decodeBody(c, &requestModel); err != nil {
		return badRequest(c, err)
	}
	consumer, consumerErr := parseAddress("consumer", requestModel.Consumer)
	if err := errors.Join(consumerErr, positive("amount", requestModel.Amount)); err != nil {
		return badRequest(c, err)
	}
	transactionInfo, err := svc.ChangeGranularFee(consumer, requestModel.Amount)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, transactionInfo)
}
