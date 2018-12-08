package sorters

import (
	"sort"

	"gitlab.com/rosenpin/git-project-showcaser/models"
)

type forks struct{}

func newForks() Sorter {
	return &forks{}
}

func (forks *forks) Sort(projects models.Projects) models.Projects {
	if projects == nil {
		panic("can't sort nil projects object")
	}

	sort.Slice(projects, func(i, j int) bool {
		return projects[i].Forks > projects[j].Forks
	})

	return projects
}

func init() {
	sortersCreators[models.ForksSort] = newForks
}
