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

func NewOracle(log *logrus.Logger) (*Oracle, error) {
	return &Oracle{log: log}, nil
}

func (d *Oracle) Withdraw(c echo.Context) error {
	var requestModel api.OracleWithdrawRequestModel
	c.Bind(&requestModel)
	transactionInfo, err := d.service.Oracle.Withdraw(requestModel.Address, requestModel.Amount)
	return c.String(http.StatusOK, transactionInfo)
}
