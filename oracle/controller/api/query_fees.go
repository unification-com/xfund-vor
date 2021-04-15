package api

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"net/http"
	"oracle/models/api"
)

func (d *Oracle) QueryFees(c echo.Context) error {
	var requestModel api.OracleQueryFeesModel
	json.NewDecoder(c.Request().Body).Decode(&requestModel)
	fee, err := d.service.QueryFees(requestModel.Consumer)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, fee)
}
