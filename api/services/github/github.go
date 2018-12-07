package services

import (
	"time"

	"gitlab.com/rosenpin/git-project-showcaser/api/models"
	"gitlab.com/rosenpin/git-project-showcaser/api/utils"
)

// Github is a git service that projects can be fetched from
type Github struct {
	*GithubFetcher
}

// NewGithub creates a new Github object and uses the timeout as the request timeout for API requests
func NewGithub(timeout time.Duration) *Github {
	return &Github{NewGithubFetcher(utils.NewHTTPJsonFetcher(timeout))}
}

func (github *Github) GetProjectsFor(username string) ([]*models.Project, error) {
	return github.FetchProjects(username)
}
