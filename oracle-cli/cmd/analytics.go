package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"oraclecli/utils"
	"strconv"

	"github.com/spf13/cobra"
)

type Prices struct {
	Eth float64 `json:"eth"`
	Usd float64 `json:"usd"`
}
type CoinGeckoResponse struct {
	Xfund Prices `json:"xfund"`
}

var (
	ifGas    uint
	ifFees   float64
	consumer string
)


// aboutCmd represents the about command
var analyticsCmd = &cobra.Command{
	Use:   "analytics [num_to_analyse]",
	Short: "Basic analytics summary",
	Long:  `
Run some basic analytics to return gas, gas price, costs and fee statistics.
Pass the number of successful requests you want to analyse, for example 100
to analyse the last 100 successful fulfillments.

Example:

oraclecli analytics 100
oraclecli analytics 1000
`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires a number to analyse argument")
		}

		numToAnalyse, err := strconv.Atoi(args[0])
		if err != nil {
			return err
		}

		if numToAnalyse <= 0 {
			return errors.New("number to analyse must be > 0")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {

		numToAnalyse, _ := strconv.Atoi(args[0])

		cgPrices := GetXfundPrice()

		url := fmt.Sprintf("%s/analytics?eth=%f&usd=%f&limit=%d&gasprice=0&fees=0&consumer=&sim=0",
			utils.OracleAddress(), cgPrices.Xfund.Eth, cgPrices.Xfund.Usd, numToAnalyse)

		client := &http.Client{}
		// Create a Bearer string by appending string access token
		var bearer = "Bearer " + utils.Settings.Settings.GetOracleKey()
		req, err := http.NewRequest("GET", url, nil)
		// add authorization header to the req
		req.Header.Add("Authorization", bearer)

		resp, err := client.Do(req)

		if err != nil {
			fmt.Println(`Sorry, something went wrong =(`)
			fmt.Println(err)
			return
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		fmt.Println(string(body))
	},
}

func GetXfundPrice() *CoinGeckoResponse {
	client := &http.Client{}
	cgUrl := "https://api.coingecko.com/api/v3/simple/price?ids=xfund&vs_currencies=eth%2Cusd"
	req, err := http.NewRequest("GET", cgUrl, nil)
	if err != nil {
		fmt.Println(err)
		return &CoinGeckoResponse{}
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return &CoinGeckoResponse{}
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		fmt.Println(readErr)
		return &CoinGeckoResponse{}
	}

	cgResp := &CoinGeckoResponse{}

	jsonErr := json.Unmarshal(body, cgResp)
	if jsonErr != nil {
		fmt.Println(jsonErr)
	}
	return cgResp
}

func init() {
	rootCmd.AddCommand(analyticsCmd)
}
