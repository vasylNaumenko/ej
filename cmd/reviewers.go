package cmd

import (
	"github.com/spf13/cobra"

	repoCfg "github.com/vasylNaumenko/ej/internal/repository/config"
	serviceCfg "github.com/vasylNaumenko/ej/internal/service/config"
)

// reviewersCmd represents the reviewers command
var reviewersCmd = &cobra.Command{
	Use:   "reviewers",
	Short: "returns a reviewers list from the configuration file",
	Long:  `returns a reviewers list from the configuration file, includes the name and the tag`,
	Run: func(cmd *cobra.Command, args []string) {
		app := getApp()

		serviceCfg.NewService(repoCfg.NewRepository(*app.Config)).
			EchoListOfAllReviewers()
	},
}

func init() {
	getCmd.AddCommand(reviewersCmd)
}
