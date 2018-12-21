package showcase

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"gitlab.com/rosenpin/git-project-showcaser/api/manager"
	"gitlab.com/rosenpin/git-project-showcaser/api/services"
	"gitlab.com/rosenpin/git-project-showcaser/api/services/github"
	"gitlab.com/rosenpin/git-project-showcaser/app"
	"gitlab.com/rosenpin/git-project-showcaser/config"
	"gitlab.com/rosenpin/git-project-showcaser/models"
	yaml "gopkg.in/yaml.v2"
)

var (
	platforms = map[string]services.ServiceCreator{
		"github": github.NewGithub,
	}
)

// ProjectsShowcase is the exported app that can be used to serve this app under an HTTP server
type ProjectsShowcase struct {
}

// NewProjectShowcase creates a new projectsShowcase object
func NewProjectShowcase() *ProjectsShowcase {
	return &ProjectsShowcase{}
}

// CreateHandler creates the HTTP handler for this project
func (projectShowcase *ProjectsShowcase) CreateHandler(configPath string) http.Handler {
	// Load config
	config := loadConfig(configPath)

	manager := manager.New(config).From(platforms[config.GitPlatform](config))

	projects, err := manager.Fetch()
	if err != nil {
		panic(err)
	}

	server := app.New(projects, config)
	ticker := time.NewTicker(config.ReloadInterval)

	go startReloading(ticker, server, manager)

	return server
}

// startReloading reloads the projects in a set interval using the ticker
func startReloading(ticker *time.Ticker, server *app.Server, manager *manager.Manager) {
	for {
		select {
		case <-ticker.C:
			reload(server, manager)
		}
	}
}

// reload reloads the projects
func reload(server *app.Server, manager *manager.Manager) {
	projects, err := manager.Fetch()
	if err != nil {
		fmt.Println("Error reloading projects: ", err)
	}

	fmt.Println("Reloaded projects")

	server.SetProjects(projects)
}

func loadConfig(configPath string) *models.Config {
	configLoader := config.NewLoader(configPath)

	config := &models.Config{}

	err := configLoader.Load(yaml.Unmarshal, config)
	if err != nil {
		panic(err)
	}

	validateConfig(config)

	return config
}

func validateConfig(config *models.Config) error {
	if config.Username == "" {
		return fmt.Errorf("invalid username")
	}

	if _, err := os.Stat(config.ResourcesPath); err != nil {
		return fmt.Errorf("invalid resources path - %v", err)
	}

	return nil
}
