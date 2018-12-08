package sorters

import (
	"gitlab.com/rosenpin/git-project-showcaser/models"
)

type alphabetically struct{}

func newalphabetically() Sorter {
	return &alphabetically{}
}

func (alphabetically *alphabetically) Sort(projects models.Projects) models.Projects {
	if projects == nil {
		panic("can't sort nil projects object")
	}

	return projects
}

func init() {
	sortersCreators[models.AlphabeticallySort] = newalphabetically
}
