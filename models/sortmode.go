package models

// SortMode is used to define the sort mode for repos
type SortMode uint8

const (
	// Stars sorts by most stars for a repo
	Stars = iota
	// Forks sorts by most forks for a repo
	Forks = iota
	// Alphabetically sorts by repo name
	Alphabetically = iota
)
