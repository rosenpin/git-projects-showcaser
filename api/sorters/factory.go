package sorters

import "gitlab.com/rosenpin/git-project-showcaser/models"

type sorterCreator func() Sorter

var (
	sortersCreators = map[models.SortMode]sorterCreator{}
)

// Create creates a new sorter object
func Create(mode models.SortMode) Sorter {
	return sortersCreators[mode]()
}
