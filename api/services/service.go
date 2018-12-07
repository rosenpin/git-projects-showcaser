package services

import "gitlab.com/rosenpin/git-project-showcaser/api/models"

// Service represents a git service that projects can be fetched from
type Service interface {
	GetProjectsFor(username string) ([]*models.Project, error)
}
