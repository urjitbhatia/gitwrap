package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of go-template-repo",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("go-template-repo version: %v, commit: %v, built at: %v\n", version, commit, date)
	},
}
