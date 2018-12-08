package filters

import (
	"gitlab.com/rosenpin/git-project-showcaser/models"
)

// Filter is the interface used to create filter objects
// Filter objects are used to filter projects, for example by max projects count, fork or not, etc.
type Filter interface {
	Filter(models.Projects) models.Projects
}

type Filters []Filter

func New(filters ...Filter) *Filters {
	this := Filters{}
	this = append(this, filters...)
	return &this
}

func (filters *Filters) Add(filter Filter) {
	*filters = append(*filters, filter)
}

func (filters *Filters) Filter(projects models.Projects) models.Projects {
	filteredProjects := projects
	for _, filter := range *filters {
		filteredProjects = filter.Filter(filteredProjects)
	}

	return filteredProjects
}
