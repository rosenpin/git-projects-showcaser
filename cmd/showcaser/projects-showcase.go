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
	yaml "gopkg.in/yaml.v2"
)

var (
	platforms = map[string]services.ServiceCreator{
		"github": github.NewGithub,
	}
)

func main() {
	// Parse flags
	var configPath string

	flag.StringVar(&configPath, "c", "", "path to the configuration file")
	flag.Parse()

	// Load config
	config := loadConfig(configPath)

	manager := manager.New(config).
		From(platforms[config.GitPlatform](config))

	projects, err := manager.Fetch()
	if err != nil {
		panic(err)
	}

	server := server.New(projects)
	go server.Start(config, err)

	ticker := time.NewTicker(config.ReloadInterval)
	for {
		select {
		case <-ticker.C:
			reload(server, manager)
		}
	}
}

func reload(server *server.Server, manager *manager.Manager) {
	projects, err := manager.Fetch()
	if err != nil {
		fmt.Println("Error reloading projects: ", err)
	}

	fmt.Println("Reloaded projects")

	server.SetProjects(projects)
}

func loadConfig(configPath string) *models.Config {
	if configPath == "" {
		panic("no configuration file specified")
	}

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
