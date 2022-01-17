package cmd

import (
	"github.com/spf13/cobra"

	"github.com/vasylNaumenko/ej/internal/service/jira_caller"
)

// usersCmd represents the users command
var usersCmd = &cobra.Command{
	Use:   "users",
	Short: "returns a list of users",
	Long:  `returns a list of users`,
	Run: func(cmd *cobra.Command, args []string) {
		app := getApp()

		jira_caller.
			NewService(app.RepoJira, app.Logger).
			GetUsers()
	},
}

func init() {
	getCmd.AddCommand(usersCmd)
}
