package parsers

import (
	"gitlab.com/rosenpin/git-project-showcaser/api/models"
)

// ProjectsParser is the interface used to parse projects data from raw json data
type ProjectsParser interface {
	Parse(interface{}) ([]*models.Project, error)
}
