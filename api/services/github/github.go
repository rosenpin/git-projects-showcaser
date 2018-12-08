package github

import (
	"time"

	"gitlab.com/rosenpin/git-project-showcaser/api/services"
	"gitlab.com/rosenpin/git-project-showcaser/api/utils"
	"gitlab.com/rosenpin/git-project-showcaser/models"
)

// Github is a git service that projects can be fetched from
type Github struct {
	*Fetcher
}

// NewGithub creates a new Github object and uses the timeout as the request timeout for API requests
func NewGithub(config *models.Config) services.Service {
	return &Github{NewFetcher(config, utils.NewHTTPJsonFetcher(time.Duration(config.HTTPRequestTimeout)*time.Second))}
}

// GetProjects gets the projects from Github and returns it
func (github *Github) GetProjects() (models.Projects, error) {
	return github.FetchProjects()
}
