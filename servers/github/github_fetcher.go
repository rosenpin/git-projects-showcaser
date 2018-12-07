package servers

import (
	"fmt"

	"gitlab.com/rosenpin/git-project-showcaser/models"
	"gitlab.com/rosenpin/git-project-showcaser/parsers"
	"gitlab.com/rosenpin/git-project-showcaser/utils"
)

const (
	githubAPIUrl = "https://api.github.com/users/%v/repos?sort=%v"
)

// GithubFetcher fetches projects from Github
type GithubFetcher struct {
	*utils.HTTPJsonFetcher
	parsers.ProjectsParser
}

// NewGithubFetcher creates a new github fetcher object used to fetch users projects
func NewGithubFetcher(fetcher *utils.HTTPJsonFetcher) *GithubFetcher {
	return &GithubFetcher{fetcher, newGithubParser()}
}

// FetchProjects is used to fetch projects of a user by his username
func (github *GithubFetcher) FetchProjects(username string) ([]*models.Project, error) {
	result, err := github.FetchJSON(fmt.Sprintf(githubAPIUrl, username, "stars"))
	if err != nil {
		return nil, err
	}

	return github.Parse(result)
}
