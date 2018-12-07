package services

import (
	"fmt"
	"testing"
	"time"

	"gitlab.com/rosenpin/git-project-showcaser/api/utils"
)

func TestGithubFetcher(t *testing.T) {
	f := NewGithubFetcher(utils.NewHTTPJsonFetcher(10 * time.Second))
	projects, err := f.FetchProjects("rosenpin")
	if err != nil {
		t.Error(err)
	}

	for _, project := range projects {
		if project == nil {
			continue
		}

		fmt.Print(*project)
	}
}
