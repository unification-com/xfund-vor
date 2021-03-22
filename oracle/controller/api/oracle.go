package api

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
	"oracle/models/api"
	"oracle/service"
)

type Oracle struct {
	log     *logrus.Logger
	service *service.Service
	context context.Context
}

func NewOracle(ctx context.Context, log *logrus.Logger, service *service.Service) (*Oracle, error) {
	return &Oracle{log: log, context: ctx, service: service}, nil
}

func (d *Oracle) Withdraw(c echo.Context) error {
	var requestModel api.OracleWithdrawRequestModel
	c.Bind(&requestModel)
	transactionInfo, err := d.service.Oracle.Withdraw(requestModel.Address, requestModel.Amount)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, transactionInfo)
	}
	return c.JSON(http.StatusOK, transactionInfo)
}
