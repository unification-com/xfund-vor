package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"net/http"
	"oraclecli/utils"
)

// aboutCmd represents the about command
var analyticsConsumersCmd = &cobra.Command{
	Use:   "consumers [consumer_address]",
	Short: "detailed consumer analytics",
	Long:  `
Run some analytics to return gas, gas price, costs and fee statistics
for all contracts served. Results are grouped by each consumer contract.

Optional [consumer_address] can be passed to filter by that address

Example:

oraclecli analytics consumers
oraclecli analytics consumers 0x1234AbcD...
`,
	Run: func(cmd *cobra.Command, args []string) {
		cgPrices := GetXfundPrice()

		c := ""
		if len(args) > 0 {
			c = args[0]
		}

		url := fmt.Sprintf("%s/consumers?eth=%f&usd=%f&consumer=%s",
			utils.OracleAddress(), cgPrices.Xfund.Eth, cgPrices.Xfund.Usd, c)

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

func init() {
	analyticsCmd.AddCommand(analyticsConsumersCmd)
}
