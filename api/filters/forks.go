package filters

import (
	"fmt"

	"gitlab.com/rosenpin/git-project-showcaser/models"
)

// Forks is a filter used to filter projects that are forks
type Forks struct {
	maxProjects uint
}

// NewForksFilter creates a new Forks filter object
func NewForksFilter(config *models.Config) Filter {
	fmt.Println()
	return &Forks{config.MaxProjects}
}

// Filter returns projects without forks
func (max *Forks) Filter(projects models.Projects) models.Projects {
	filteredProjects := models.Projects{}
	for _, project := range projects {
		if project.IsFork {
			continue
		}
		filteredProjects = append(filteredProjects, project)
	}
	return filteredProjects
}

func init() {
	filterCreators[models.ForksFilter] = NewForksFilter
}
