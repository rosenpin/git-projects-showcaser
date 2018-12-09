package server

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
}

// New creates a new server object
func New(projects models.Projects) *Server {
	return &Server{projects}
}

// SetProjects sets the projects that the server will return, this is used to update the server without having to restart it
func (server *Server) SetProjects(projects models.Projects) {
	server.projects = projects
}

// Start starts the HTTP server on the specified port
func (server *Server) Start(config *models.Config, err error) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if err != nil {
			fmt.Fprintf(w, fmt.Sprint("Error: ", err))
			return
		}

		file, err := ioutil.ReadFile(path.Join(config.ResourcesPath, "index.html"))
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
		}{config.Username, server.projects, config.GitPlatform, config.ProfileURL}

		err = template.Execute(w, page)
		if err != nil {
			fmt.Fprintf(w, fmt.Sprint("Error: ", err))
			return
		}
	})

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(path.Join(config.ResourcesPath, "static")))))

	if err := http.ListenAndServe(fmt.Sprintf(":%d", uint(config.Port)), nil); err != nil {
		panic(err)
	}
}
