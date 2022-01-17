package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/vasylNaumenko/ej/internal/domain/config"
	"github.com/vasylNaumenko/ej/internal/service/jira_caller"
	"github.com/vasylNaumenko/ej/internal/service/notifier"
)

const separator = ","

var (
	assignees   string // list of assignee`s tags
	randomCount int
)

// reviewCmd represents the review command
var reviewCmd = &cobra.Command{
	Use:   "review",
	Short: "Creates tasks for review of your MR.",
	Long: `Creates tasks for review of yours merge request.
	Examples:
		review [issue-id] [MR link] -t=tag1,tag2 (creates tasks for the tag1 and the tag2 assignees)
		review [issue-id] [MR link] -t=tag1 -r=1  (creates tasks for the tag1 assignee and plus a random one)	
		review [issue-id] [MR link] -r=2  (creates tasks for the 2 random reviewers)
`,
	Run: func(cmd *cobra.Command, args []string) {
		app := getApp()

		// validation
		if len(args) < 2 {
			err := fmt.Errorf(`please, specify the command arguments. Use review -h for details
`)
			exitIfError(err)
		}

		if assignees == "" && randomCount == 0 {
			err := fmt.Errorf(`the tags or the random count must be specified
	use review -h for details
`)
			exitIfError(err)
		}

		// Add account owner to the list that excludes from the random choosing of reviewers.
		mySelf, err := app.Config.Reviewers.GetByTag(app.Config.MyAccountTag)
		if err != nil {
			err := fmt.Errorf("can`t find a MyAccount tag")
			exitIfError(err)
		}
		excludeList := config.Assignees{
			mySelf,
		}

		// select assignees by tags
		var assigneeList config.Assignees
		tags := strings.Split(assignees, separator)
		for _, tag := range tags {
			reviewer, err := app.Config.Reviewers.GetByTag(tag)
			if err != nil {
				exitIfError(err)
			}
			assigneeList = append(assigneeList, reviewer)
			excludeList = append(excludeList, reviewer)
		}

		// Select or adding a random assignees.
		if randomCount > 0 {
			assigneeListRandom, err := app.Config.Reviewers.GetRandom(2, excludeList...)
			exitIfError(err)

			assigneeList = append(assigneeList, assigneeListRandom...)
		}

		// Creating the tasks and generating the links.
		ids := assigneeList.GetSliceOfIDs()
		links, err := jira_caller.
			NewService(app.RepoJira, app.Logger).
			CreateReviewIssue(args[0], args[1], ids)
		exitIfError(err)

		err = notifier.NewService(app.RepoNotifier).Notify(assigneeList, links)
		exitIfError(err)

		// echo to console
		for _, recipient := range assigneeList {
			fmt.Printf("Notification was sent for: %s\n", recipient.Name)
		}
	},
}

func init() {
	rootCmd.AddCommand(reviewCmd)

	reviewCmd.Flags().StringVarP(&assignees,
		"tags", "t", "",
		"assignees tags: -t=[tag1,tag2,...]")
	reviewCmd.Flags().IntVarP(&randomCount,
		"random", "r", 0,
		"random assignees count: -r=2 (takes random 2 assignee from a reviewers list)")
}
