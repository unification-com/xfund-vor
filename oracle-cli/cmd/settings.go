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
	"fmt"
	"oraclecli/utils"
	"syscall"

	"github.com/spf13/cobra"
)

// settingsCmd represents the settings command
var settingsCmd = &cobra.Command{
	Use:   "settings",
	Short: "Oracle CLI settings command",
	Long:  `Use this command to setup your CLI (or change existing settings).`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print(utils.Settings.String())
		_, err := GetSettings()
		if err != nil {
			fmt.Println(err)
		}
	},
}

func GetSettings() (input string, err error) {
	fmt.Println("")
	fmt.Println("Okay, let's configure your CLI")
	fmt.Println("What do you want to do?")
	fmt.Println("1 - Set HTTP/cli Oracle Key")
	fmt.Println("2 - Set Oracle host address")
	fmt.Println("3 - Set Oracle host port")
	fmt.Println("0 - Exit")
	fmt.Print("Action: ")
	_, err = fmt.Scanf("%s\n", &input)

	switch input {
	case "1":
		key, err := GetOracleKey()
		if err != nil {
			fmt.Println(err)
		}
		err = utils.Settings.SetOracleKey(key)
		if err != nil {
			fmt.Println(err)
		}
	case "2":
		key, err := GetOracleHost()
		if err != nil {
			fmt.Println(err)
		}
		err = utils.Settings.SetOracleHost(key)
		if err != nil {
			fmt.Println(err)
		}
	case "3":
		key, err := GetOraclePort()
		if err != nil {
			fmt.Println(err)
		}
		err = utils.Settings.SetOraclePort(key)
		if err != nil {
			fmt.Println(err)
		}
	case "0":
		syscall.Exit(0)
		return

	default:
		fmt.Println("Please choose action.")
		input, err = GetSettings()
	}
	input, err = GetSettings()
	return
}

func GetOracleHost() (input string, err error) {
	fmt.Println("")
	fmt.Print("Oracle host address (ex.: 127.0.0.1): ")
	_, err = fmt.Scanf("%s\n", &input)
	if input == "" {
		fmt.Println("Please enter Oracle host address.")
		input, err = GetOracleHost()
	}
	return
}

func GetOraclePort() (input string, err error) {
	fmt.Println("")
	fmt.Print("Oracle host port (ex.: 8888): ")
	_, err = fmt.Scanf("%s\n", &input)
	if input == "" {
		fmt.Println("Please enter Oracle host port.")
		input, err = GetOraclePort()
	}
	return
}

func GetOracleKey() (input string, err error) {
	fmt.Println("")
	fmt.Print("Oracle host key (ex.: rod0gbc3mhyxdiah2vwialx1q3osk5cw): ")
	_, err = fmt.Scanf("%s\n", &input)
	if input == "" {
		fmt.Println("Please enter Oracle host key.")
		input, err = GetPrivateKey()
	}
	return
}

func init() {
	rootCmd.AddCommand(settingsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// settingsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// settingsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
