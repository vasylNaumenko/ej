package jira

import (
	"fmt"
	"io"

	"github.com/andygrunwald/go-jira"

	domain "ej/internal/domain/config"
)

type (
	// Repository defines a repository for jira API.
	Repository struct {
		config domain.Config
		client *jira.Client
	}
)

// NewRepository constructor
func NewRepository(config domain.Config) *Repository {
	tp := jira.BasicAuthTransport{
		Username: config.User,
		Password: config.Token,
	}
	jiraClient, err := jira.NewClient(tp.Client(), config.BaseUrl)
	if err != nil {
		panic(err)
	}

	return &Repository{
		config: config,
		client: jiraClient,
	}
}

// GetMyProfile returns profile name
func (r Repository) GetMyProfile() (string, error) {
	getSelf, _, err := r.client.User.GetSelf()
	if err != nil {
		return "", err
	}

	return getSelf.DisplayName, nil
}

// GetProjects returns list of projects
func (r Repository) GetProjects() ([]string, error) {
	list, _, err := r.client.Project.GetList()
	if err != nil {
		return nil, err
	}

	var res []string
	res = append(res, "ID: Name: Key")
	for _, s := range *list {
		res = append(res, fmt.Sprintf("%s: %s: %s", s.ID, s.Name, s.Key))
	}

	return res, nil
}

// GetUsers returns list of active atlassian users
func (r Repository) GetUsers() ([]string, error) {
	const pageSize = 50
	var res []string
	res = append(res, "ID: Name")
	counter := 0
	for {
		users, _, err := r.client.User.Find(".",
			jira.WithActive(true),
			jira.WithStartAt(counter),
			jira.WithMaxResults(pageSize),
		)
		if err != nil {
			return nil, err
		}

		if len(users) == 0 {
			break
		}

		for _, s := range users {
			if s.AccountType != "atlassian" {
				continue
			}
			res = append(res, fmt.Sprintf("%s: %s", s.AccountID, s.DisplayName))
		}

		counter += pageSize
	}

	return res, nil
}

// CreateIssue create a new issue for review issueID
func (r Repository) CreateIssue(issueKey, assigneeID, description string) (string, error) {
	issue, _, err := r.client.Issue.Get(issueKey, nil)
	if err != nil {
		return "", err
	}

	newIssue, resp, err := r.client.Issue.Create(
		getNewIssue(issue.Key, assigneeID, description),
	)
	if err != nil {
		return "", getErrorFromBody(resp)
	}

	resp, err = r.client.Issue.AddLink(
		getNewLink(newIssue, issue),
	)
	if err != nil {
		return "", getErrorFromBody(resp)
	}

	return r.config.BaseUrl + issueBrowsePath + newIssue.Key, nil
}

// getErrorFromBody puts a response body in an error
func getErrorFromBody(resp *jira.Response) error {
	defer resp.Body.Close()
	b, _ := io.ReadAll(resp.Body)
	return fmt.Errorf("%s", string(b))
}
