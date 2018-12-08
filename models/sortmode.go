package models

// SortMode is used to define the sort mode for repos
type SortMode uint8

const (
	// Stars sorts by most stars for a repo
	Stars SortMode = iota
	// Forks sorts by most forks for a repo
	Forks SortMode = iota
	// Alphabetically sorts by repo name
	Alphabetically SortMode = iota
)
