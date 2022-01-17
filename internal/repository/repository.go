package repository

import "ej/internal/domain/config"

type (
	Jira interface {
		GetMyProfile() (string, error)
		CreateIssue(issueID, assigneeID, description string) (string, error)
		GetProjects() ([]string, error)
		GetUsers() ([]string, error)
	}

	Config interface {
		GetAllReviewers() []string
	}

	Notifier interface {
		Notify(recipients config.Assignees, links map[string]string) error
	}
)
