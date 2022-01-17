package jira

import (
	"github.com/andygrunwald/go-jira"
	"github.com/trivago/tgo/tcontainer"
)

// Some predefined constants, used for issue creation and linking.
const (
	summaryText     = "Ревью на"
	priority        = "2"
	projectID       = "10016"
	issueTypeName   = "Task"
	labelsNames     = "review"
	linkType        = "review"
	issueBrowsePath = "/browse/"
)

// getNewIssue returns an issue with a prepared data.
func getNewIssue(issueKey, assigneeID, description string) *jira.Issue {
	summary := summaryText + " " + issueKey

	issue := &jira.Issue{
		Fields: &jira.IssueFields{
			Summary: summary,
			Type: jira.IssueType{
				Name: issueTypeName,
			},
			Project: jira.Project{
				ID: projectID,
			},
			Assignee: &jira.User{
				AccountID: assigneeID,
			},
			Priority:    &jira.Priority{ID: priority}, // high
			Description: description,
			Labels:      []string{labelsNames}, // review
		},
	}

	setCustomField(issue)

	return issue
}

// setCustomField sets a Department custom field.
func setCustomField(issue *jira.Issue) {
	customFields := tcontainer.NewMarshalMap()
	customFields.Set("customfield_10027", []map[string]string{{"id": "10020"}})
	issue.Fields.Unknowns = customFields
}

// getNewLink returns a new link for the parent and child issues.
func getNewLink(newIssue *jira.Issue, issue *jira.Issue) *jira.IssueLink {
	return &jira.IssueLink{
		Type: jira.IssueLinkType{Name: linkType},
		InwardIssue: &jira.Issue{
			ID: newIssue.ID,
		},
		OutwardIssue: &jira.Issue{
			ID: issue.ID,
		},
	}
}
