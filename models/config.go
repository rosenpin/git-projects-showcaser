package models

import "time"

// Config is the structs that holds the app configurable data
type Config struct {
	ResourcesPath      string        `json:"ResourcesPath" yaml:"ResourcesPath"`
	Port               float64       `json:"Port" yaml:"Port"`
	Username           string        `json:"Username" yaml:"Username"`
	AuthCode           string        `json:"AuthCode,omitempty" yaml:"AuthCode"`
	HTTPRequestTimeout time.Duration `json:"HTTPRequestTimeout" yaml:"HTTPRequestTimeout"`
	MaxProjects        uint          `json:"MaxProjects" yaml:"MaxProjects"`
	IncludeForks       bool          `json:"IncludeForks" yaml:"IncludeForks"`
	GitPlatform        string        `json:"GitPlatform" yaml:"GitPlatform"`
	SortMode           string        `json:"SortMode" yaml:"SortMode"`
	ProfileURL         string        `json:"ProfileURL" yaml:"ProfileURL"`
}
