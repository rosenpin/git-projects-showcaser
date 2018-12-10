package app

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"path"

	"gitlab.com/rosenpin/git-project-showcaser/models"
)

// Server is the object responsible of the HTTP server
type Server struct {
	projects models.Projects
	config   *models.Config
}

// New creates a new server object
func New(projects models.Projects, config *models.Config) *Server {
	return &Server{projects, config}
}

// SetProjects sets the projects that the server will return, this is used to update the server without having to restart it
func (server *Server) SetProjects(projects models.Projects) {
	server.projects = projects
}

func (server *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	file, err := ioutil.ReadFile(path.Join(server.config.ResourcesPath, "index.html"))
	if err != nil {
		fmt.Fprintf(w, fmt.Sprint("Error: ", err))
		return
	}

	template, err := template.New("Web Page").Parse(string(file))
	if err != nil {
		fmt.Fprintf(w, fmt.Sprint("Error: ", err))
		return
	}

	page := struct {
		Title       string
		Projects    models.Projects
		ServiceName string
		ProfileURL  string
	}{server.config.Username, server.projects, server.config.GitPlatform, server.config.ProfileURL}

	err = template.Execute(w, page)
	if err != nil {
		fmt.Fprintf(w, fmt.Sprint("Error: ", err))
		return
	}
}
