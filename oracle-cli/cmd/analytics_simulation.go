package cmd

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"oraclecli/utils"
	"strconv"

	"github.com/spf13/cobra"
)

// aboutCmd represents the about command
var analyticsSimCmd = &cobra.Command{
	Use:   "sim [num_to_analyse]",
	Short: "simulate analytics with given params",
	Long:  `
Pass simulation values for gas price and fees using the --if-gas and --if-fees
flags respectively. These will be applied to the analytics results in place of
the real values to see what costs/fees would be like with these simulated values.

use the [num_to_analyse] arg to limit the number of requests to analyse.

The --consumer flag can be used to filter a single consumer contract, which can help
set granular fees for individual consumers.

Example:

oraclecli analytics sim 1000 -g 150 -f 0.05
oraclecli analytics sim 200 --if-gas=150 --if-fees=0.1
oraclecli analytics sim 200 --if-gas=100 --if-fees=0.05 --consumer=0x232AbC...
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

		if ifGas == 0 {
			return errors.New("enter if-gas value")
		}
		if ifFees == 0 {
			return errors.New("enter if-fees value")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {

		numToAnalyse, _ := strconv.Atoi(args[0])

		cgPrices := GetXfundPrice()

		url := fmt.Sprintf("%s/analytics?eth=%f&usd=%f&limit=%d&gasprice=%d&fees=%f&sim=1&consumer=%s",
			utils.OracleAddress(), cgPrices.Xfund.Eth, cgPrices.Xfund.Usd, numToAnalyse, ifGas, ifFees, consumer)

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
	analyticsSimCmd.Flags().StringVarP(&consumer, "consumer", "q", "", "filter by consumer contract address")
	analyticsSimCmd.Flags().UintVarP(&ifGas, "if-gas", "g", 0, "manually set gas price in Gwei for analytics to simulate costs at this price")
	analyticsSimCmd.Flags().Float64VarP(&ifFees, "if-fees", "f", 0, "manually set XFUND fees for analytics to simulate costs at this price")

	analyticsCmd.AddCommand(analyticsSimCmd)
}
