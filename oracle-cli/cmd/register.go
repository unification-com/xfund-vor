/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"net/http"
	"oraclecli/models"
	"oraclecli/utils"
	"strconv"
)

// registerCmd represents the register command
var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "Register your new Oracle",
	Long: `Use this command to register your new oracle
Usage:
 oracle-cli register [account name] [0x... private key] [fee] [provider pays gas (true/false)]
`,
	Run: func(cmd *cobra.Command, args []string) {
		fee, err := strconv.ParseInt(args[2], 10, 64)
		if err != nil {
			fmt.Println("Incorrect amount")
			return
		}
		paysgas, err := strconv.ParseBool(args[3])
		if err != nil {
			fmt.Println("Incorrect provider pays gas parameter")
			return
		}
		requestStruct := models.OracleRegisterRequestModel{
			AccountName:     args[0],
			PrivateKey:      args[1],
			Fee:             fee,
			ProviderPaysGas: paysgas,
		}
		requestJSON, err := json.Marshal(requestStruct)
		if err != nil {
			fmt.Println("Can't marshal request")
			return
		}
		request := bytes.NewBuffer(requestJSON)
		resp, err := http.Post(fmt.Sprint(utils.OracleAddress(), "/register"), "encoding/json", request)
		if err != nil {
			fmt.Println("Something went wrong.")
		}
		fmt.Println(resp)
	},
}

func init() {
	rootCmd.AddCommand(registerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// registerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// registerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
