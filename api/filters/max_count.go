package filters

import "gitlab.com/rosenpin/git-project-showcaser/models"

const (
	// MaxCountDefault is the default value of the filter, if set, the filter will do nothing
	MaxCountDefault = 0
)

// MaxCount is a filter used to limit the number of projects returned
type MaxCount struct {
	maxProjects uint
}

// NewMaxFilter creates a new MaxCount filter object
func NewMaxFilter(config *models.Config) Filter {
	return &MaxCount{config.MaxProjects}
}

// Filter returns the first n projects from the projects slice
func (max *MaxCount) Filter(projects models.Projects) models.Projects {
	if max.maxProjects == MaxCountDefault {
		return projects
	}

	return projects[:max.maxProjects]
}

func init() {
	filterCreators[models.MaxFilter] = NewMaxFilter
}
