package sorters

import "gitlab.com/rosenpin/git-project-showcaser/models"

// Sorter is used to sort projects
type Sorter interface {
	Sort(models.Projects) models.Projects
}
