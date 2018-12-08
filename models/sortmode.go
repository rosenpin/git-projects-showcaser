package models

// SortMode is used to define the sort mode for repos
type SortMode uint8

var (
	// SortFromConfig is used to get the correct sort mode from the configuration string
	SortFromConfig = map[string]SortMode{
		"stars":          StarsSort,
		"forks":          ForksSort,
		"alphabetically": AlphabeticallySort,
	}
)

const (
	// StarsSort sorts by most stars for a repo
	StarsSort SortMode = iota
	// ForksSort sorts by most forks for a repo
	ForksSort SortMode = iota
	// AlphabeticallySort sorts by repo name
	AlphabeticallySort SortMode = iota
)
