package fetchers

import "gitlab.com/rosenpin/git-project-showcaser/models"

// Fetcher fetches projects from the appropriate git server
type Fetcher interface {
	FetchProjects(username string) ([]models.Project, error)
}
