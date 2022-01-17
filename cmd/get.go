package cmd

import (
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "get [projects, reviewers, users]",
	Long:  `gets list by a given command`,
}

func init() {
	rootCmd.AddCommand(getCmd)
}
