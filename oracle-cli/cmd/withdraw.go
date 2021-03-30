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

// withdrawCmd represents the withdraw command
var withdrawCmd = &cobra.Command{
	Use:   "withdraw",
	Short: "Withdraw your xFUND",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		amount, err := GetAmount()
		address, err := GetAddress()
		requestStruct := models.OracleWithdrawRequestModel{
			Address: address,
			Amount:  amount,
		}
		requestJSON, err := json.Marshal(requestStruct)
		if err != nil {
			fmt.Println("Can't marshal request")
			return
		}
		request := bytes.NewBuffer(requestJSON)

		// Create a Bearer string by appending string access token
		var bearer = "Bearer " + utils.Settings.Settings.GetOracleKey()
		req, err := http.NewRequest("POST", fmt.Sprint(utils.OracleAddress(), "/withdraw"), request)
		// add authorization header to the req
		req.Header.Add("Authorization", bearer)
		client := &http.Client{}
		resp, err := client.Do(req)

		if err != nil {
			fmt.Println("Something went wrong.")
		}
		fmt.Println(resp)
	},
}

func GetAmount() (input int64, err error) {
	var rawInput string
	fmt.Println("")
	fmt.Print("Amount: ")
	_, err = fmt.Scanf("%s\n", &rawInput)
	input, err = strconv.ParseInt(rawInput, 10, 64)
	if err != nil {
		fmt.Println("Incorrect amount")
		input, err = GetAmount()
	}
	if input == 0 {
		fmt.Println("Please enter Fee.")
		input, err = GetAmount()
	}
	return
}

func GetAddress() (input string, err error) {
	fmt.Println("")
	fmt.Print("Address: ")
	_, err = fmt.Scanf("%s\n", &input)
	if input == "" {
		fmt.Println("Please enter address.")
		input, err = GetPrivateKey()
	}
	return
}

func init() {
	rootCmd.AddCommand(withdrawCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// withdrawCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// withdrawCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
