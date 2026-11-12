package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"oracle/config"
	controller "oracle/controller/api"
	"oracle/service"
)

// The admin API routes each request to the chain it names: by name or network id, required when the oracle
// serves several chains, and 503 for a chain whose worker isn't running yet.
func TestAdminAPIRoutesByChain(t *testing.T) {
	path := filepath.Join(t.TempDir(), "config.json")
	const coordinator = `"contract_address":"0xCeB679DE5fEF4F7c2a0e4cB6A4a0C6e6d5A50ED8"`
	require.NoError(t, os.WriteFile(path, []byte(`{"chains":[
		{"name":"polygon","network_id":137,"eth_http_host":"http://127.0.0.1:1",`+coordinator+`},
		{"name":"mainnet","network_id":1,"eth_http_host":"http://127.0.0.1:1",`+coordinator+`}]}`), 0600))
	conf, err := config.NewConfig(path)
	require.NoError(t, err)
	original := config.Conf
	config.Conf = conf
	t.Cleanup(func() { config.Conf = original })

	workers := newSupervisor(context.Background(), nil, nil, 0)
	polygon := &service.Service{Chain: conf.ChainList()[0]}
	workers.services[137] = polygon // mainnet hasn't started

	api, err := controller.NewOracle(context.Background(), log, workers.service)
	require.NoError(t, err)
	var served *service.Service
	route := api.Chain(func(c echo.Context, svc *service.Service) error {
		served = svc
		return c.String(http.StatusOK, svc.Chain.Name)
	})

	for selector, want := range map[string]int{"polygon": 200, "137": 200, "mainnet": 503, "1": 503, "bsc": 400, "": 400} {
		served = nil
		rec := httptest.NewRecorder()
		require.NoError(t, route(echo.New().NewContext(httptest.NewRequest(http.MethodGet, "/about?chain="+selector, nil), rec)))
		assert.Equal(t, want, rec.Code, "chain=%q: %s", selector, rec.Body.String())
		if want == 200 {
			assert.Same(t, polygon, served)
		} else {
			assert.Nil(t, served, "chain=%q must not reach the handler", selector)
		}
	}
}
