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
	"net/http"
	"oraclecli/models"
	"oraclecli/utils"
	"strconv"

	"github.com/spf13/cobra"
)

// setproviderpaysgasCmd represents the setproviderpaysgas command
var setproviderpaysgasCmd = &cobra.Command{
	Use:   "setproviderpaysgas",
	Short: "Set who pays gas",
	Long: `A command to set who pays gas - true if provider pays gas

Usage:
 oracle-cli setproviderpaysgas [true/false]
`,
	Run: func(cmd *cobra.Command, args []string) {
		paysgas, err := strconv.ParseBool(args[3])
		if err != nil {
			fmt.Println("Incorrect provider pays gas parameter")
			return
		}
		requestStruct := models.OracleSetProviderPaysGasRequestModel{
			ProviderPays: paysgas,
		}
		requestJSON, err := json.Marshal(requestStruct)
		if err != nil {
			fmt.Println("Can't marshal request")
			return
		}
		request := bytes.NewBuffer(requestJSON)
		resp, err := http.Post(fmt.Sprint(utils.OracleAddress(), "/setproviderpaysgas"), "encoding/json", request)
		if err != nil {
			fmt.Println("Something went wrong.")
		}
		fmt.Println(resp)
	},
}

func init() {
	rootCmd.AddCommand(setproviderpaysgasCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// setproviderpaysgasCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// setproviderpaysgasCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
