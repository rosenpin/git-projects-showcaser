package services

import (
	"time"

	"gitlab.com/rosenpin/git-project-showcaser/models"
)

// Service represents a git service that projects can be fetched from
type Service interface {
	GetProjectsFor(username string) ([]*models.Project, error)
}

// ServiceCreator is the function used to create services (services constructors)
type ServiceCreator func(timeout time.Duration) Service
