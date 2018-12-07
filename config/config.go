package config

import (
	"encoding/json"
	"io/ioutil"

	"gitlab.com/rosenpin/git-project-showcaser/models"
)

// Loader loads the configuration from the configuration file into the configuration object
type Loader struct{ path string }

// NewLoader creates a new configuration loader object using the provided path
func NewLoader(path string) *Loader {
	return &Loader{path}
}

// Load loads the configuration from the configuration file into the configuration obejct
func (cl *Loader) Load(target *models.Config) error {
	content, err := ioutil.ReadFile(cl.path)
	if err != nil {
		return err
	}

	return json.Unmarshal(content, target)
}
