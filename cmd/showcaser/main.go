package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"gitlab.com/rosenpin/git-project-showcaser/api/manager"
	"gitlab.com/rosenpin/git-project-showcaser/api/services"
	"gitlab.com/rosenpin/git-project-showcaser/api/services/github"
	"gitlab.com/rosenpin/git-project-showcaser/config"
	"gitlab.com/rosenpin/git-project-showcaser/models"
	"gitlab.com/rosenpin/git-project-showcaser/server"
)

var (
	platforms = map[string]services.ServiceCreator{
		"github": github.NewGithub,
	}

	sortModes = map[string]models.SortMode{
		"stars":          models.Stars,
		"forks":          models.Forks,
		"alphabetically": models.Alphabetically,
	}
)

const (
	defaultTimeout = 10 * time.Second
)

func main() {
	// Parse flags
	var gitPlatform, sortMode, configPath string
	var maxProjects int
	var includeForks bool

	flag.StringVar(&gitPlatform, "git", "github", "git platform to use")
	flag.StringVar(&sortMode, "sort", "stars", "projects sort mode")
	flag.StringVar(&configPath, "c", "", "path to the configuration file")
	flag.IntVar(&maxProjects, "max", 10, "max number of projects to fetch")
	flag.BoolVar(&includeForks, "forks", false, "include forks")
	flag.Parse()

	// Load config
	config := loadConfig(configPath)

	manager := manager.New().
		From(platforms[gitPlatform](defaultTimeout)).
		UsingSortMode(sortModes[sortMode]).
		WithNoMoreThan(maxProjects).
		ForUser(config.Username)

	if includeForks {
		manager = manager.IncludingForks()
	}

	projects, err := manager.Fetch()
	server.StartServer(config, projects, err)
}

func loadConfig(configPath string) *models.Config {
	if configPath == "" {
		panic("no configuration file specified")
	}

	configLoader := config.NewLoader(configPath)

	config := &models.Config{}

	err := configLoader.Load(config)
	if err != nil {
		panic(err)
	}

	validateConfig(config)

	return config
}

func validateConfig(config *models.Config) error {
	if config.Port > 65535 || config.Port <= 0 {
		return fmt.Errorf("invalid port number")
	}

	if config.Username == "" {
		return fmt.Errorf("invalid username")
	}

	if _, err := os.Stat(config.ResourcesPath); err != nil {
		return fmt.Errorf("invalid resources path - %v", err)
	}

	return nil
}
