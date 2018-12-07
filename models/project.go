package models

// Project is the structs that holds data about a project
type Project struct {
	Name        string  `github:"name"`
	Description string  `github:"description"`
	Link        string  `github:"html_url"`
	Language    string  `github:"language"`
	IsFork      bool    `github:"fork"`
	Stars       float64 `github:"stargazers_count"`
	Forks       float64 `github:"forks"`
}

// NewProject creates a new project object using the project data
func NewProject(
	isFork bool,
	stars float64,
	forks float64,
	name string,
	description string,
	link string,
	language string,
) *Project {
	return &Project{
		IsFork:      isFork,
		Stars:       stars,
		Forks:       forks,
		Name:        name,
		Description: description,
		Link:        link,
		Language:    language}
}
