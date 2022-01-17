package jira_caller

import (
	"fmt"

	"github.com/rs/zerolog"

	"ej/internal/repository"
)

// Service defines the jira_caller service
type Service struct {
	log      *zerolog.Logger
	repoJira repository.Jira
}

func NewService(repoJira repository.Jira, log *zerolog.Logger) *Service {
	return &Service{
		log:      log,
		repoJira: repoJira,
	}
}

// GetProjects echoes a jira projects list into the console
func (s Service) GetProjects() {
	list, err := s.repoJira.GetProjects()
	if err != nil {
		s.log.Error().Err(err).Msg("Cant get projects")
		return
	}

	for _, info := range list {
		fmt.Println(info)
	}

}

// GetUsers echoes a jira users list into the console
func (s Service) GetUsers() {
	list, err := s.repoJira.GetUsers()
	if err != nil {
		s.log.Error().Err(err).Msg("Cant get projects")
		return
	}

	for _, info := range list {
		fmt.Println(info)
	}
}

// CreateReviewIssue creates an issue for each of the assignee and returns a links list for the new issues,
func (s Service) CreateReviewIssue(issueID, description string, assigneeIDs []string) (map[string]string, error) {
	var links = make(map[string]string, len(assigneeIDs))

	for _, assigneeID := range assigneeIDs {
		link, err := s.repoJira.CreateIssue(issueID, assigneeID, description)
		if err != nil {
			s.log.Error().Err(err).Msg("Can`t create an issue")
			return nil, err
		}
		links[assigneeID] = link
	}

	return links, nil
}
