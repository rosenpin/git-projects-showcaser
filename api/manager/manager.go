package manager

import (
	"gitlab.com/rosenpin/git-project-showcaser/api/models"
	"gitlab.com/rosenpin/git-project-showcaser/api/services"
)

type Filters struct {
	sortMode    *models.SortMode
	maxProjects uint
	withForks   bool
}

type Manager struct {
	service  services.Service
	filters  *Filters
	username string
}

func New() *Manager {
	return &Manager{}
}

func (manager *Manager) clear() {
	manager.filters = nil
	manager.username = ""
}

func (manager *Manager) From(service services.Service) *Manager {
	manager.service = service
	return manager
}

func (manager *Manager) UsingSortMode(sortmode *models.SortMode) *Manager {
	manager.filters.sortMode = sortmode
	return manager
}

func (manager *Manager) WithNoMoreThan(maxProjects uint) *Manager {
	manager.filters.maxProjects = maxProjects
	return manager
}

func (manager *Manager) WithForks() *Manager {
	manager.filters.withForks = true
	return manager
}

func (manager *Manager) WithoutForks() *Manager {
	manager.filters.withForks = false
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
