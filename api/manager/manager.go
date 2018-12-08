package manager

import (
	filtering "gitlab.com/rosenpin/git-project-showcaser/api/filters"
	"gitlab.com/rosenpin/git-project-showcaser/api/services"
	"gitlab.com/rosenpin/git-project-showcaser/api/sorters"
	"gitlab.com/rosenpin/git-project-showcaser/models"
)

type Manager struct {
	service services.Service
}

func New(config *models.Config) *Manager {
	return &Manager{}
}

func (manager *Manager) From(service services.Service) *Manager {
	manager.service = service
	return manager
}

func (manager *Manager) Fetch(config *models.Config) (models.Projects, error) {
	projects, err := manager.service.GetProjects()
	if err != nil {
		return nil, err
	}

	filters := filtering.New()
	if !config.IncludeForks {
		filters.Add(filtering.Create(models.ForksFilter, config))
	}
	filters.Add(filtering.Create(models.MaxFilter, config))

	sorter := sorters.Create(models.SortFromConfig[config.SortMode])

	return filters.Filter(sorter.Sort(projects)), nil
}
