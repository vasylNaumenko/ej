package config

import (
	"fmt"

	"ej/internal/repository"
)

// Service defines the config service
type Service struct {
	repoConfig repository.Config
}

func NewService(repoConfig repository.Config) *Service {
	return &Service{
		repoConfig: repoConfig,
	}
}

func (s Service) EchoListOfAllReviewers() {
	list := s.repoConfig.GetAllReviewers()
	for _, line := range list {
		fmt.Println(line)
	}
}
