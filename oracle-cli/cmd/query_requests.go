package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"net/http"
	"oraclecli/utils"
)

var (
	page   uint
	limit  uint
	status int
	order  string
)

// aboutCmd represents the about command
var queryrequestsCmd = &cobra.Command{
	Use:   "queryrequests",
	Short: "get the requests from the DB",
	Long:  `Query for all paginated randomness requests that match optional filters.

Status filters:
The --status | -s flag accepts the following:

 -1 = no filter, return all results (default)
 0  = unknown status
 1  = request initialised
 2  = fulfill tx sent
 3  = failed to fulfil
 4  = fulfillment succeeded

Examples:
$ oraclecli queryrequests --page=2 --limit=20
$ oraclecli queryrequests --status=3
$ oraclecli queryrequests --order=asc
`,
	Run: func(cmd *cobra.Command, args []string) {

		// Create a Bearer string by appending string access token
		var bearer = "Bearer " + utils.Settings.Settings.GetOracleKey()
		url := fmt.Sprintf("%s/requests?page=%d&limit=%d&order=%s&status=%d", utils.OracleAddress(), page, limit, order, status)
		fmt.Println("url", url)
		req, err := http.NewRequest("GET", url, nil)
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
		fmt.Println(string(body))
	},
}

func init() {
	queryrequestsCmd.Flags().UintVarP(&page, "page", "p", 1, "page number")
	queryrequestsCmd.Flags().UintVarP(&limit, "limit", "l", 10, "results to return per page")
	queryrequestsCmd.Flags().IntVarP(&status, "status", "s", -1, "request status")
	queryrequestsCmd.Flags().StringVarP(&order, "order", "o", "desc", "order asc | desc")
	rootCmd.AddCommand(queryrequestsCmd)
}

