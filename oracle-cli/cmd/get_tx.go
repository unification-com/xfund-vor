package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"net/http"
	"oraclecli/utils"
)

// aboutCmd represents the about command
var gettxCmd = &cobra.Command{
	Use:   "gettx",
	Short: "get tx info",
	Long:  `Returns tx info, including the receipt if available`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires a tx hash")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {

		txHash := args[0]

		url := fmt.Sprintf("%s/tx?tx_hash=%s", utils.OracleAddress(), txHash)

		// Create a Bearer string by appending string access token
		var bearer = "Bearer " + utils.Settings.Settings.GetOracleKey()
		req, err := http.NewRequest("GET", url, nil)
		// add authorization header to the req
		req.Header.Add("Authorization", bearer)
		client := &http.Client{}
		resp, err := client.Do(req)

		if err != nil {
			fmt.Println("Something went wrong.")
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		fmt.Println(string(body))
	},
}

func init() {
	rootCmd.AddCommand(gettxCmd)
}
