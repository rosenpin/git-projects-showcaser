package services

import (
	"gitlab.com/rosenpin/git-project-showcaser/models"
)

// Service represents a git service that projects can be fetched from
type Service interface {
	// GetProjects returns the projects from the appropriate service
	GetProjects() ([]*models.Project, error)
}

// ServiceCreator is the function used to create services (services constructors)
type ServiceCreator func(config *models.Config) Service
