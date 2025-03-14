package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "datamarket",
	Short: "A data trading market",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("data market")
	},
}

func Exceute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(DaemonCmd)
	rootCmd.AddCommand(versionCmd)
}
