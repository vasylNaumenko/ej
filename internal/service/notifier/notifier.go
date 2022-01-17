package notifier

import (
	"ej/internal/domain/config"
	"ej/internal/repository"
)

// Service defines the notifier service
type Service struct {
	repoNotifier repository.Notifier
}

func NewService(repoNotifier repository.Notifier) *Service {
	return &Service{repoNotifier: repoNotifier}
}

// Notify calls the notify method from a Notifier repository.
func (s Service) Notify(recipients config.Assignees, links map[string]string) error {
	return s.repoNotifier.Notify(recipients, links)
}
