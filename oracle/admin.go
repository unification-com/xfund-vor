package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math"
	"net"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"

	"oracle/config"
	"oracle/models/api"
)

// apiClient calls a running oracle's HTTP API: its address comes from the oracle's config, its key from the
// file `oracle api-key` wrote (api_key beside the config, unless --api-key names another).
type apiClient struct {
	base string
	key  string
	http *http.Client
}

func newAPIClient() (*apiClient, error) {
	keyFile := options.APIKeyFile
	if keyFile == "" {
		keyFile = defaultAPIKeyFile()
	}
	key, err := os.ReadFile(keyFile)
	if err != nil {
		return nil, fmt.Errorf("reading the API key (issue one with `oracle api-key`): %w", err)
	}
	host := config.Conf.Serve.Host
	if ip := net.ParseIP(host); host == "" || (ip != nil && ip.IsUnspecified()) {
		host = "127.0.0.1"
	}
	return &apiClient{
		base: "http://" + net.JoinHostPort(host, strconv.Itoa(int(config.Conf.Serve.Port))),
		key:  strings.TrimSpace(string(key)),
		http: &http.Client{Timeout: 2 * time.Minute},
	}, nil
}

// call sends method path?query with body as JSON (if any) and returns the response body. A status other than
// 2xx is an error carrying the body, which holds the oracle's reason.
func (c *apiClient) call(method, path string, query url.Values, body interface{}) ([]byte, error) {
	var reader io.Reader
	if body != nil {
		encoded, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		reader = bytes.NewReader(encoded)
	}
	target := c.base + path
	if len(query) > 0 {
		target += "?" + query.Encode()
	}
	req, err := http.NewRequest(method, target, reader)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+c.key)
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("is the oracle running? %w", err)
	}
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return nil, fmt.Errorf("%s %s: %s: %s", method, path, resp.Status, strings.TrimSpace(string(respBody)))
	}
	return respBody, nil
}

// adminCommand is a command that calls the running oracle: run gets a client, and prints what it returns.
func adminCommand(use, short string, args cobra.PositionalArgs, run func(c *apiClient, args []string) (string, error)) *cobra.Command {
	cmd := &cobra.Command{
		Use:   use,
		Short: short,
		Args:  args,
		RunE: withConfig(func(_ *cobra.Command, args []string) error {
			client, err := newAPIClient()
			if err != nil {
				return err
			}
			out, err := run(client, args)
			if err != nil {
				return err
			}
			fmt.Println(strings.TrimSpace(out))
			return nil
		}),
	}
	cmd.Flags().StringVar(&options.APIKeyFile, "api-key", "", "path to the file holding the API key (default: api_key beside the config)")
	return cmd
}

// raw passes the response through as text.
func raw(b []byte, err error) (string, error) {
	return string(b), err
}

func parseAmount(name, s string) (int64, error) {
	n, err := strconv.ParseInt(s, 10, 64)
	if err != nil || n <= 0 {
		return 0, fmt.Errorf("%s must be a whole number above zero, in the token's smallest unit", name)
	}
	return n, nil
}

// xfund converts an amount in xFUND's smallest unit (9 decimals) for display.
func xfund(tokens string) string {
	n, err := strconv.ParseInt(strings.Trim(strings.TrimSpace(tokens), `"`), 10, 64)
	if err != nil {
		return strings.TrimSpace(tokens)
	}
	return fmt.Sprintf("%d (%f xFUND)", n, float64(n)/math.Pow10(9))
}

// xfundPrice is xFUND's price in ETH and USD from CoinGecko, for the analytics' cost figures (zero if
// unavailable: the oracle then reports costs in gas only).
func xfundPrice() (eth, usd float64) {
	var price struct {
		Xfund struct {
			Eth float64 `json:"eth"`
			Usd float64 `json:"usd"`
		} `json:"xfund"`
	}
	resp, err := (&http.Client{Timeout: 15 * time.Second}).Get("https://api.coingecko.com/api/v3/simple/price?ids=xfund&vs_currencies=eth%2Cusd")
	if err != nil {
		fmt.Fprintln(os.Stderr, "xFUND price unavailable:", err)
		return 0, 0
	}
	defer resp.Body.Close()
	if err = json.NewDecoder(resp.Body).Decode(&price); err != nil {
		fmt.Fprintln(os.Stderr, "xFUND price unavailable:", err)
	}
	return price.Xfund.Eth, price.Xfund.Usd
}

func adminCommands() []*cobra.Command {
	status := adminCommand("status", "Check the oracle is running", cobra.NoArgs, func(c *apiClient, _ []string) (string, error) {
		return raw(c.call(http.MethodGet, "/status", nil, nil))
	})
	about := adminCommand("about", "Show the oracle's address, key hash, balances and fee", cobra.NoArgs, func(c *apiClient, _ []string) (string, error) {
		return raw(c.call(http.MethodGet, "/about", nil, nil))
	})
	stop := adminCommand("stop", "Stop the oracle", cobra.NoArgs, func(c *apiClient, _ []string) (string, error) {
		return raw(c.call(http.MethodPost, "/stop", nil, nil))
	})
	withdraw := adminCommand("withdraw <amount> <address>", "Withdraw earned xFUND (amount in its smallest unit) to an address", cobra.ExactArgs(2), func(c *apiClient, args []string) (string, error) {
		amount, err := parseAmount("amount", args[0])
		if err != nil {
			return "", err
		}
		return raw(c.call(http.MethodPost, "/withdraw", nil, api.OracleWithdrawRequestModel{Address: args[1], Amount: amount}))
	})
	changeFee := adminCommand("changefee <amount>", "Change the oracle's fee (in xFUND's smallest unit)", cobra.ExactArgs(1), func(c *apiClient, args []string) (string, error) {
		amount, err := parseAmount("fee", args[0])
		if err != nil {
			return "", err
		}
		return raw(c.call(http.MethodPost, "/changefee", nil, api.OracleChangeFeeRequestModel{Amount: amount}))
	})
	changeGranularFee := adminCommand("changegranularfee <amount> <consumer>", "Change the fee one consumer contract pays", cobra.ExactArgs(2), func(c *apiClient, args []string) (string, error) {
		amount, err := parseAmount("fee", args[0])
		if err != nil {
			return "", err
		}
		return raw(c.call(http.MethodPost, "/changegranularfee", nil, api.OracleChangeGranularFeeRequestModel{Consumer: args[1], Amount: amount}))
	})
	queryFees := adminCommand("queryfees [consumer]", "Show the oracle's fee, or the one a consumer contract pays", cobra.MaximumNArgs(1), func(c *apiClient, args []string) (string, error) {
		var consumer string
		if len(args) > 0 {
			consumer = args[0]
		}
		fee, err := c.call(http.MethodPost, "/queryfees", nil, api.OracleQueryFeesModel{Consumer: consumer})
		if err != nil {
			return "", err
		}
		return "Fee: " + xfund(string(fee)), nil
	})
	queryWithdrawable := adminCommand("querywithdrawable", "Show the xFUND the oracle can withdraw", cobra.NoArgs, func(c *apiClient, _ []string) (string, error) {
		tokens, err := c.call(http.MethodGet, "/querywithdrawable", nil, nil)
		if err != nil {
			return "", err
		}
		return "Withdrawable tokens: " + xfund(string(tokens)), nil
	})
	getTx := adminCommand("gettx <tx_hash>", "Show a transaction and its receipt", cobra.ExactArgs(1), func(c *apiClient, args []string) (string, error) {
		return raw(c.call(http.MethodGet, "/tx", url.Values{"tx_hash": {args[0]}}, nil))
	})

	var page, limit, requestStatus int
	var order string
	queryRequests := adminCommand("queryrequests", "List randomness requests", cobra.NoArgs, func(c *apiClient, _ []string) (string, error) {
		return raw(c.call(http.MethodGet, "/requests", url.Values{
			"page": {strconv.Itoa(page)}, "limit": {strconv.Itoa(limit)}, "status": {strconv.Itoa(requestStatus)}, "order": {order},
		}, nil))
	})
	queryRequests.Flags().IntVarP(&page, "page", "p", 1, "page number")
	queryRequests.Flags().IntVarP(&limit, "limit", "l", 10, "requests per page (at most 100)")
	queryRequests.Flags().IntVarP(&requestStatus, "status", "s", -1, "only requests with this status (-1: any)")
	queryRequests.Flags().StringVarP(&order, "order", "o", "desc", "asc or desc")

	register := adminCommand("register <account> <fee>", "Add a proving key to the keystore and register it with the coordinator (the key is asked for)", cobra.ExactArgs(2), func(c *apiClient, args []string) (string, error) {
		fee, err := parseAmount("fee", args[1])
		if err != nil {
			return "", err
		}
		if !stdinIsTerminal() {
			return "", fmt.Errorf("the private key is asked for: %w", errNoTerminal)
		}
		privateKey, err := promptSecret("Private key (hex; not shown as you type): ")
		if err != nil {
			return "", err
		}
		return raw(c.call(http.MethodPost, "/register", nil, api.OracleRegisterRequestModel{AccountName: args[0], PrivateKey: privateKey, Fee: fee}))
	})

	analyticsArgs := func(_ *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("requires the number of fulfilled requests to analyse")
		}
		_, err := parseAmount("the number to analyse", args[0])
		return err
	}
	analytics := adminCommand("analytics <num_to_analyse>", "Gas, cost and fee statistics for the last fulfilled requests", analyticsArgs, func(c *apiClient, args []string) (string, error) {
		eth, usd := xfundPrice()
		return raw(c.call(http.MethodGet, "/analytics", url.Values{
			"eth": {fmt.Sprint(eth)}, "usd": {fmt.Sprint(usd)}, "limit": {args[0]}, "gasprice": {"0"}, "fees": {"0"}, "consumer": {""}, "sim": {"0"},
		}, nil))
	})
	var ifGas uint
	var ifFees float64
	var simConsumer string
	sim := adminCommand("sim <num_to_analyse>", "Analytics at a simulated gas price and fee", analyticsArgs, func(c *apiClient, args []string) (string, error) {
		if ifGas == 0 || ifFees == 0 {
			return "", errors.New("--if-gas and --if-fees are required")
		}
		eth, usd := xfundPrice()
		return raw(c.call(http.MethodGet, "/analytics", url.Values{
			"eth": {fmt.Sprint(eth)}, "usd": {fmt.Sprint(usd)}, "limit": {args[0]}, "gasprice": {fmt.Sprint(ifGas)},
			"fees": {fmt.Sprint(ifFees)}, "consumer": {simConsumer}, "sim": {"1"},
		}, nil))
	})
	sim.Flags().UintVarP(&ifGas, "if-gas", "g", 0, "gas price in gwei to simulate")
	sim.Flags().Float64VarP(&ifFees, "if-fees", "f", 0, "fee in xFUND to simulate")
	sim.Flags().StringVarP(&simConsumer, "consumer", "q", "", "only this consumer contract")
	consumers := adminCommand("consumers [consumer]", "Per-consumer statistics", cobra.MaximumNArgs(1), func(c *apiClient, args []string) (string, error) {
		var consumer string
		if len(args) > 0 {
			consumer = args[0]
		}
		eth, usd := xfundPrice()
		return raw(c.call(http.MethodGet, "/consumers", url.Values{"eth": {fmt.Sprint(eth)}, "usd": {fmt.Sprint(usd)}, "consumer": {consumer}}, nil))
	})
	analytics.AddCommand(sim, consumers)

	return []*cobra.Command{status, about, stop, withdraw, changeFee, changeGranularFee, queryFees, queryWithdrawable, getTx, queryRequests, register, analytics}
}
