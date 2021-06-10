package cmd

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"oraclecli/utils"
	"strings"

	"github.com/spf13/cobra"
)

// aboutCmd represents the about command
var querywithdrawableCmd = &cobra.Command{
	Use:   "querywithdrawable",
	Short: "get the amount of xFUND you can withdraw",
	Long:  `Returns the amount of withdrawable xFUND currently held for your Oracle in the VORCoordinator smart contract`,
	Run: func(cmd *cobra.Command, args []string) {

		// Create a Bearer string by appending string access token
		var bearer = "Bearer " + utils.Settings.Settings.GetOracleKey()
		req, err := http.NewRequest("GET", utils.OracleAddress()+"/querywithdrawable", nil)
		// add authorization header to the req
		req.Header.Add("Authorization", bearer)
		client := &http.Client{}
		resp, err := client.Do(req)

		if err != nil {
			fmt.Println(`Sorry, something went wrong =(`)
			fmt.Println(err)
			return
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		tokens := strings.TrimSpace(string(body))
		xfund := utils.ConvertToXfund(tokens)
		fmt.Println("Withdrawable Tokens:", tokens, fmt.Sprintf("(%f xFUND)", xfund))
	},
}

func init() {
	rootCmd.AddCommand(querywithdrawableCmd)
}
