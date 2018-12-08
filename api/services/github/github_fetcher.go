package github

import (
	"fmt"

	"gitlab.com/rosenpin/git-project-showcaser/api/parsers"
	"gitlab.com/rosenpin/git-project-showcaser/api/utils"
	"gitlab.com/rosenpin/git-project-showcaser/models"
)

const (
	githubAPIUrl = "https://api.github.com/users/%v/repos?per_page=50?access_token=%v"
)

// Fetcher fetches projects from Github
type Fetcher struct {
	*utils.HTTPJsonFetcher
	parsers.ProjectsParser
	config *models.Config
}

// NewFetcher creates a new github fetcher object used to fetch users projects
func NewFetcher(config *models.Config, fetcher *utils.HTTPJsonFetcher) *Fetcher {
	return &Fetcher{fetcher, newGithubParser(), config}
}

// FetchProjects is used to fetch projects of a user by his username
func (github *Fetcher) FetchProjects() (models.Projects, error) {
	apiURL := fmt.Sprintf(githubAPIUrl, github.config.Username, github.config.AuthCode)
	fmt.Println("querying: ", apiURL)
	result, err := github.FetchJSON(apiURL)
	if err != nil {
		return nil, err
	}

	return github.Parse(result)
}
