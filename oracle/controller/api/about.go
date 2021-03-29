package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (d *Oracle) About(c echo.Context) error {
	info, err := d.service.About()
	if err != nil {
		return c.String(http.StatusInternalServerError, info)
	}
	return c.String(http.StatusOK, info)
}
