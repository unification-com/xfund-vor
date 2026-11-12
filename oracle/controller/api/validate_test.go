package api

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"oracle/service"
)

// Bad input must be refused before it reaches the service: the Oracle here has none, so anything that got
// through would panic.
func TestAdminHandlersRefuseBadInput(t *testing.T) {
	d := &Oracle{}
	const address = "0x90F8bf6A479f320ead074411a4B0e7944Ea8c9C1"
	cases := []struct {
		name    string
		handler func(echo.Context, *service.Service) error
		body    string
	}{
		{"withdraw: not JSON", d.Withdraw, `not json`},
		{"withdraw: empty", d.Withdraw, `{}`},
		{"withdraw: to 0x0", d.Withdraw, `{"address":"0x0000000000000000000000000000000000000000","amount":1}`},
		{"withdraw: no 0x", d.Withdraw, `{"address":"90F8bf6A479f320ead074411a4B0e7944Ea8c9C1","amount":1}`},
		{"withdraw: zero amount", d.Withdraw, `{"address":"` + address + `","amount":0}`},
		{"change fee: zero", d.ChangeFee, `{"amount":0}`},
		{"change fee: negative", d.ChangeFee, `{"amount":-5}`},
		{"granular fee: bad consumer", d.ChangeGranularFee, `{"consumer":"0x1","amount":5}`},
		{"granular fee: zero", d.ChangeGranularFee, `{"consumer":"` + address + `","amount":0}`},
		{"register: no account", d.Register, `{"private_key":"0x4f3edf983ac636a65a842ce7c78d9aa706d3b113bce9c46f30d7d21715b23b1d","fee":1}`},
		{"register: bad key", d.Register, `{"account_name":"spare","private_key":"0xsecretish","fee":1}`},
		{"register: zero fee", d.Register, `{"account_name":"spare","private_key":"0x4f3edf983ac636a65a842ce7c78d9aa706d3b113bce9c46f30d7d21715b23b1d","fee":0}`},
		{"query fees: bad consumer", d.QueryFees, `{"consumer":"nope"}`},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			ctx := echo.New().NewContext(httptest.NewRequest(http.MethodPost, "/", strings.NewReader(c.body)), rec)
			require.NoError(t, c.handler(ctx, nil))
			assert.Equal(t, http.StatusBadRequest, rec.Code, rec.Body.String())
			assert.NotContains(t, rec.Body.String(), "secretish", "errors must not echo a private key")
		})
	}
}
