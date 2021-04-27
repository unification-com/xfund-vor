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
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"oraclecli/utils"
)

var cfgFile string

func initSettings() {
	fmt.Println("cfgFile", cfgFile)
	settings, err := utils.NewSettingsStore(cfgFile)
	if err != nil {
		fmt.Println(err)
	}
	utils.Settings = settings
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "oraclecli",
	Short: "A CLI to manage your Oracle",
	Long: `CLI to manage your Oracle.

Note:
 You need to run "oracle start -c [config_path | optional] -k [key | optional]" to start your daemon before using CLI.
`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		initSettings()
	},
	Run: func(cmd *cobra.Command, args []string) {
		// Create a Bearer string by appending string access token
		var bearer = "Bearer " + utils.Settings.Settings.GetOracleKey()
		req, err := http.NewRequest("GET", utils.OracleAddress()+"/status", nil)
		// add authorization header to the req
		req.Header.Add("Authorization", bearer)
		client := &http.Client{}
		resp, err := client.Do(req)

		if err != nil {
			fmt.Println(`Sorry, something went wrong =(`)
			return
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		fmt.Println("Oracle status: ", string(body))
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {

	homepath, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	rootCmd.PersistentFlags().StringVarP(&cfgFile, "conf", "c",
		filepath.Join(homepath, ".oracle-cli_settings.json"), "oraclecli settings file")

	cobra.CheckErr(rootCmd.Execute())
}
