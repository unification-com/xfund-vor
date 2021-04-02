package api

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"net/http"
	"oracle/models/api"
)

func (d *Oracle) ChangeFee(c echo.Context) error {
	var requestModel api.OracleChangeFeeRequestModel
	json.NewDecoder(c.Request().Body).Decode(&requestModel)
	transactionInfo, err := d.service.ChangeFee(requestModel.Amount)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, transactionInfo)
}
