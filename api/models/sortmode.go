package models

// SortMode is used to define the sort mode for repos
type SortMode uint8

const (
	stars          = iota
	forks          = iota
	alphabetically = iota
)
