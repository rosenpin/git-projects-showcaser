package github

import (
	"fmt"

	"gitlab.com/rosenpin/git-project-showcaser/api/parsers"
	"gitlab.com/rosenpin/git-project-showcaser/models"
)

const (
	githubTag = "github"
)

var (
	errInvalidMessageFormat = fmt.Errorf("invalid message format")
)

type githubParser struct{}

func newGithubParser() *githubParser {
	return &githubParser{}
}

func (github *githubParser) Parse(raw interface{}) ([]*models.Project, error) {
	rawProjects, ok := raw.([]interface{})
	if !ok {
		return nil, errInvalidMessageFormat
	}

	var projects []*models.Project

	for _, rawProject := range rawProjects {
		mappedProject := rawProject.(map[string]interface{})

		project, err := parsers.CreateProjectUsingTags(mappedProject, githubTag)
		if err != nil {
			return nil, err
		}

		projects = append(projects, project)
	}

	return projects, nil
}
