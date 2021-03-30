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
`,
	Run: func(cmd *cobra.Command, args []string) {
		err := Register(cmd, args)
		if err != nil {
			fmt.Println(err)
		}
	},
}

func Register(cmd *cobra.Command, args []string) (err error) {
	accountName, err := GetUsername()
	privateKey, err := GetPrivateKey()
	fee, err := GetFee()
	paysGas, err := GetProviderPaysGas()
	if err != nil {
		fmt.Println("Sorry, there is a problem with data you entered =(")
		err = Register(cmd, args)
		return
	}
	requestStruct := models.OracleRegisterRequestModel{
		AccountName:     accountName,
		PrivateKey:      privateKey,
		Fee:             fee,
		ProviderPaysGas: paysGas,
	}
	requestJSON, err := json.Marshal(requestStruct)
	if err != nil {
		fmt.Println("Can't marshal request")
		return
	}
	request := bytes.NewBuffer(requestJSON)

	// Create a Bearer string by appending string access token
	var bearer = "Bearer " + utils.Settings.Settings.GetOracleKey()
	req, err := http.NewRequest("POST", fmt.Sprint(utils.OracleAddress(), "/register"), request)
	// add authorization header to the req
	req.Header.Add("Authorization", bearer)
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Something went wrong.")
	}
	fmt.Println(resp)
	return
}

func GetUsername() (input string, err error) {
	fmt.Println("")
	fmt.Print("Username: ")
	_, err = fmt.Scanf("%s\n", &input)
	if input == "" {
		fmt.Println("Please enter account username.")
		input, err = GetUsername()
	}
	return
}

func GetPrivateKey() (input string, err error) {
	fmt.Println("")
	fmt.Print("Private Key: ")
	_, err = fmt.Scanf("%s\n", &input)
	if input == "" {
		fmt.Println("Please enter account Private Key.")
		input, err = GetPrivateKey()
	}
	return
}

func GetFee() (input int64, err error) {
	var rawInput string
	fmt.Println("")
	fmt.Print("Fee: ")
	_, err = fmt.Scanf("%s\n", &rawInput)
	input, err = strconv.ParseInt(rawInput, 10, 64)
	if err != nil {
		fmt.Println("Incorrect amount")
		input, err = GetFee()
	}
	if input == 0 {
		fmt.Println("Please enter Fee.")
		input, err = GetFee()
	}
	return
}

func GetProviderPaysGas() (input bool, err error) {
	var rawInput string
	fmt.Println("")
	fmt.Print("Does provider pay gas? [true|false]: ")
	_, err = fmt.Scanf("%s\n", &rawInput)
	input, err = strconv.ParseBool(rawInput)
	if err != nil {
		fmt.Println("Incorrect provider pays gas parameter")
		input, err = GetProviderPaysGas()
	}

	return
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
