package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var version = "0.0.1"
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("data market version is " + version)
	},
}
