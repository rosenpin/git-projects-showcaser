package models

// FilterMode is used to define the sort mode for repos
type FilterMode uint8

const (
	// MaxFilter filters projects by projects count
	MaxFilter FilterMode = iota
	// ForksFilter filters projects that aren't forks
	ForksFilter FilterMode = iota
)
