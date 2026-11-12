package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"oracle/service"
)

func (d *Oracle) About(c echo.Context, svc *service.Service) error {
	info, err := svc.About()
	if err != nil {
		return c.String(http.StatusInternalServerError, info)
	}
	return c.String(http.StatusOK, info)
}
