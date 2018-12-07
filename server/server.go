package server

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"time"

	"gitlab.com/rosenpin/git-project-showcaser/api/models"
	servers "gitlab.com/rosenpin/git-project-showcaser/api/servers/github"
	"gitlab.com/rosenpin/git-project-showcaser/api/utils"
)

// StartServer starts the HTTP server on the specified port
func StartServer(port uint) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		githubFetcher := servers.NewGithubFetcher(utils.NewHTTPJsonFetcher(10 * time.Second))
		projects, err := githubFetcher.FetchProjects("rosenpin")
		if err != nil {
			fmt.Fprintf(w, fmt.Sprint("Error: ", err))
			return
		}
		file, err := ioutil.ReadFile("../resources/index.html")
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
		}{"title", projects}

		err = template.Execute(w, page)
		if err != nil {
			fmt.Fprintf(w, fmt.Sprint("Error: ", err))
			return
		}
	})

	http.HandleFunc("/static/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Requested ", r.RequestURI)
		content, err := ioutil.ReadFile(fmt.Sprint("../resources/", r.RequestURI))
		if err != nil {
			fmt.Fprintf(w, fmt.Sprint("error: ", err))
			return
		}

		fmt.Fprintf(w, string(content))
	})
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
