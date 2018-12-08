package sorters

import (
	"sort"

	"gitlab.com/rosenpin/git-project-showcaser/models"
)

type Stars struct{}

func NewStars() Sorter {
	return &Stars{}
}

func (stars *Stars) Sort(projects models.Projects) models.Projects {
	if projects == nil {
		panic("can't sort nil projects object")
	}

	sort.Slice(projects, func(i, j int) bool {
		return projects[i].Stars > projects[j].Stars
	})

	return projects
}

func init() {
	sortersCreators[models.StarsSort] = NewStars
}
