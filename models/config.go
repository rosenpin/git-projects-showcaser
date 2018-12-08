package models

// Config is the structs that holds the app configurable data
type Config struct {
	ResourcesPath      string  `json:"ResourcesPath"`
	Port               float64 `json:"Port"`
	Username           string  `json:"Username"`
	AuthCode           string  `json:"AuthCode,omitempty"`
	HTTPRequestTimeout int     `json:"HTTPRequestTimeout"`
}
