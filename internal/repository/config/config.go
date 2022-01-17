package config

import (
	"fmt"

	domain "ej/internal/domain/config"
)

type (
	// Repository defines a main struct for Config repository.
	Repository struct {
		config domain.Config
	}
)

func NewRepository(config domain.Config) *Repository {
	return &Repository{config: config}
}

// GetAllReviewers returns an info for all reviewers.
func (r Repository) GetAllReviewers() []string {
	var res []string

	res = append(res, "Name: Tag")
	res = append(res, fmt.Sprintf("Reviewers total: %v", len(r.config.Reviewers)))
	for _, reviewer := range r.config.Reviewers {
		res = append(res, fmt.Sprintf("\t %s: %s", reviewer.Name, reviewer.Tag))
	}

	return res
}
