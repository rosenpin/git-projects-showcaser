package manager

import (
	"gitlab.com/rosenpin/git-project-showcaser/api/services"
	"gitlab.com/rosenpin/git-project-showcaser/models"
)

type Filters struct {
	sortMode    models.SortMode
	maxProjects int
	withForks   bool
}

type Manager struct {
	service  services.Service
	filters  Filters
	username string
}

func New() *Manager {
	return &Manager{filters: Filters{}}
}

func (manager *Manager) clear() {
	manager.filters.sortMode = models.Stars
	manager.filters.maxProjects = 10
	manager.filters.withForks = false
	manager.username = ""
}

func (manager *Manager) From(service services.Service) *Manager {
	manager.service = service
	return manager
}

func (manager *Manager) UsingSortMode(sortmode models.SortMode) *Manager {
	manager.filters.sortMode = sortmode
	return manager
}

func (manager *Manager) WithNoMoreThan(maxProjects int) *Manager {
	manager.filters.maxProjects = maxProjects
	return manager
}

func (manager *Manager) IncludingForks() *Manager {
	manager.filters.withForks = true
	return manager
}

func (manager *Manager) ForUser(username string) *Manager {
	manager.username = username
	return manager
}

func (manager *Manager) Fetch() ([]*models.Project, error) {
	defer manager.clear()
	projects, err := manager.service.GetProjectsFor(manager.username)
	if err != nil {
		return nil, err
	}
	//TODO filter and sort projects

	return projects, nil
}
