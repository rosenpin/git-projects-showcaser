package server

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"path"

	"gitlab.com/rosenpin/git-project-showcaser/models"
)

// StartServer starts the HTTP server on the specified port
func StartServer(config *models.Config, projects []*models.Project, err error) {
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

		template, err := template.New("mypage").Parse(string(file))
		if err != nil {
			fmt.Fprintf(w, fmt.Sprint("Error: ", err))
			return
		}

		page := struct {
			Title    string
			Projects []*models.Project
		}{config.Username, projects}

		err = template.Execute(w, page)
		if err != nil {
			fmt.Fprintf(w, fmt.Sprint("Error: ", err))
			return
		}
	})

	http.HandleFunc("/static/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Requested ", r.RequestURI)
		content, err := ioutil.ReadFile(path.Join(config.ResourcesPath, r.RequestURI))
		if err != nil {
			fmt.Fprintf(w, fmt.Sprint("error: ", err))
			return
		}

		fmt.Fprintf(w, string(content))
	})
	http.ListenAndServe(fmt.Sprintf(":%d", uint(config.Port)), nil)
}
