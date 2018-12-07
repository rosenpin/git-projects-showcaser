package servers

import (
	"fmt"
	"testing"
	"time"

	"gitlab.com/rosenpin/git-project-showcaser/utils"
)

func TestGithubFetcher(t *testing.T) {
	f := NewGithubFetcher(utils.NewHTTPJsonFetcher(10 * time.Second))
	raw, err := f.FetchProjects("rosenpin")
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%T", raw)
}
