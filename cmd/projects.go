package cmd

import (
	"github.com/spf13/cobra"

	"ej/internal/service/jira_caller"
)

// projectsCmd represents the projects command
var projectsCmd = &cobra.Command{
	Use:   "projects",
	Short: "returns a list of projects",
	Long:  `returns a list of projects`,
	Run: func(cmd *cobra.Command, args []string) {
		app := getApp()

		jira_caller.
			NewService(app.RepoJira, app.Logger).
			GetProjects()
	},
}

func init() {
	getCmd.AddCommand(projectsCmd)
}
