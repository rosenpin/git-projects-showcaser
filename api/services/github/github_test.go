package github

import (
	"fmt"
	"testing"
	"time"

	"gitlab.com/rosenpin/git-project-showcaser/api/utils"
	"gitlab.com/rosenpin/git-project-showcaser/models"
)

func TestGithubFetcher(t *testing.T) {
	config := &models.Config{ResourcesPath: "", Port: 0, Username: "rosenpin", AuthCode: "", HTTPRequestTimeout: 10}
	f := NewFetcher(config, utils.NewHTTPJsonFetcher(10*time.Second))
	projects, err := f.FetchProjects()
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
