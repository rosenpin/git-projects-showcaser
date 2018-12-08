package models

import "time"

// Config is the structs that holds the app configurable data
type Config struct {
	ResourcesPath      string        `yaml:"ResourcesPath"`
	Port               float64       `yaml:"Port"`
	Username           string        `yaml:"Username"`
	AuthCode           string        `yaml:"AuthCode,omitempty"`
	HTTPRequestTimeout time.Duration `yaml:"HTTPRequestTimeout"`
	MaxProjects        uint          `yaml:"MaxProjects"`
	IncludeForks       bool          `yaml:"IncludeForks"`
	GitPlatform        string        `yaml:"GitPlatform"`
	SortMode           string        `yaml:"SortMode"`
	ProfileURL         string        `yaml:"ProfileURL"`
	ReloadInterval     time.Duration `yaml:"ReloadInterval"`
}
