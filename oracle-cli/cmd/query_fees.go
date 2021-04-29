package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"net/http"
	"oraclecli/models"
	"oraclecli/utils"
	"strings"
)

// aboutCmd represents the about command
var queryfeesCmd = &cobra.Command{
	Use:   "queryfees",
	Short: "get the fee",
	Long:  `Returns your current fee. Optionally pass a consumer contract address for granular fees`,
	Run: func(cmd *cobra.Command, args []string) {

		c := ""
		if len(args) > 0 {
			c = args[0]
			fmt.Println("Consumer contract:", c)
		}

		requestStruct := models.OracleQueryFeesModel{
			Consumer: c,
		}
		requestJSON, err := json.Marshal(requestStruct)
		if err != nil {
			fmt.Println("Can't marshal request")
			return
		}
		request := bytes.NewBuffer(requestJSON)

		// Create a Bearer string by appending string access token
		var bearer = "Bearer " + utils.Settings.Settings.GetOracleKey()
		req, err := http.NewRequest("POST", fmt.Sprint(utils.OracleAddress(), "/queryfees"), request)
		// add authorization header to the req
		req.Header.Add("Authorization", bearer)
		client := &http.Client{}
		resp, err := client.Do(req)

		if err != nil {
			fmt.Println("Something went wrong.")
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		tokens := strings.TrimSpace(string(body))
		xfund := utils.ConvertToXfund(tokens)
		fmt.Println("Fee:", tokens, fmt.Sprintf("(%f xFUND)", xfund))
	},
}

func init() {
	rootCmd.AddCommand(queryfeesCmd)
}

