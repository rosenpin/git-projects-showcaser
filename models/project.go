package models

// Project is the structs that holds data about a project
type Project struct {
	name        string
	description string
	link        string
	language    string
	isFork      bool
	stars       float64
	forks       float64
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
	return &Project{isFork: isFork,
		stars:       stars,
		forks:       forks,
		name:        name,
		description: description,
		link:        link,
		language:    language}
}
