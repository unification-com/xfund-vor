package api

import (
	"context"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"

	"oracle/service"
)

// ErrChainNotRunning: the chain asked for is configured, but its worker isn't running (yet).
var ErrChainNotRunning = errors.New("the chain isn't running yet")

// Oracle serves the admin API. Each request names its chain with ?chain=<name or network id>, which may be
// left out when the oracle serves one chain.
type Oracle struct {
	log      *logrus.Logger
	services func(selector string) (*service.Service, error)
	context  context.Context
}

// NewOracle serves the admin API for the chains services finds.
func NewOracle(ctx context.Context, log *logrus.Logger, services func(selector string) (*service.Service, error)) (*Oracle, error) {
	return &Oracle{log: log, context: ctx, services: services}, nil
}

// Chain adapts a handler for one chain's service into a route: it finds the chain the request names, 400 if
// there's no such chain (or it must be named), 503 if it isn't running.
func (d *Oracle) Chain(handler func(c echo.Context, svc *service.Service) error) echo.HandlerFunc {
	return func(c echo.Context) error {
		svc, err := d.services(c.QueryParam("chain"))
		if errors.Is(err, ErrChainNotRunning) {
			return c.String(http.StatusServiceUnavailable, err.Error())
		}
		if err != nil {
			return badRequest(c, err)
		}
		return handler(c, svc)
	}
}
