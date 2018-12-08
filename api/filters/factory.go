package filters

import "gitlab.com/rosenpin/git-project-showcaser/models"

type filterCreator func(*models.Config) Filter

var (
	filterCreators = map[models.FilterMode]filterCreator{}
)

// Create creates a new filter object using the configuration
func Create(mode models.FilterMode, config *models.Config) Filter {
	return filterCreators[mode](config)
}
