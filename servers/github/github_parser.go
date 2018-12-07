package servers

import (
	"fmt"

	"gitlab.com/rosenpin/git-project-showcaser/models"
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
		isFork, _ := mappedProject["fork"].(bool)
		stars, _ := mappedProject["stargazers_count"].(float64)
		forks, _ := mappedProject["forks"].(float64)
		name, _ := mappedProject["name"].(string)
		description, _ := mappedProject["description"].(string)
		url, _ := mappedProject["html_url"].(string)
		lang, _ := mappedProject["language"].(string)

		project := models.NewProject(isFork, stars, forks, name, description, url, lang)
		projects = append(projects, project)
	}

	return projects, nil
}
