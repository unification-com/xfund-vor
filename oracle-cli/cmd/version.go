package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"oraclecli/version"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "output version info",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		versionInfo := version.NewInfo()
		fmt.Println(versionInfo.String())
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

