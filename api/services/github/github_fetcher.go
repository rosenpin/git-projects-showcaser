package github

import (
	"fmt"

	"gitlab.com/rosenpin/git-project-showcaser/api/parsers"
	"gitlab.com/rosenpin/git-project-showcaser/api/utils"
	"gitlab.com/rosenpin/git-project-showcaser/models"
)

const (
	githubAPIUrl = "https://api.github.com/users/%v/repos?access_token=%v"
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
	apiURL := fmt.Sprintf(githubAPIUrl, username, "e2e933cd8f9945e19c9ad2c8c6f02782ae70f8e3")
	fmt.Println("querying: ", apiURL)
	result, err := github.FetchJSON(apiURL)
	if err != nil {
		return nil, err
	}

	return github.Parse(result)
}
