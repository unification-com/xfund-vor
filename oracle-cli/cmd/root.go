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
	"github.com/spf13/cobra"
	"io/ioutil"
	"net/http"
	"oraclecli/utils"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "oracle-cli",
	Short: "A CLI to manage your Oracle",
	Long: `CLI to manage your Oracle.

Note:
 You need to run "oracled start -c [config_path | optional] -k [key | optional]" to start your daemon before using CLI.
`,
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
	cobra.CheckErr(rootCmd.Execute())
}

//func init() {
//	cobra.OnInitialize(initConfig)
//
//	// Here you will define your flags and configuration settings.
//	// Cobra supports persistent flags, which, if defined here,
//	// will be global for your application.
//
//	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.oracle-cli.yaml)")
//
//	// Cobra also supports local flags, which will only run
//	// when this action is called directly.
//	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
//}
//
//// initConfig reads in config file and ENV variables if set.
//func initConfig() {
//	if cfgFile != "" {
//		// Use config file from the flag.
//		viper.SetConfigFile(cfgFile)
//	} else {
//		// Find home directory.
//		home, err := homedir.Dir()
//		cobra.CheckErr(err)
//
//		// Search config in home directory with name ".oracle-cli" (without extension).
//		viper.AddConfigPath(home)
//		viper.SetConfigName(".oracle-cli")
//	}
//
//	viper.AutomaticEnv() // read in environment variables that match
//
//	// If a config file is found, read it in.
//	if err := viper.ReadInConfig(); err == nil {
//		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
//	}
//}
