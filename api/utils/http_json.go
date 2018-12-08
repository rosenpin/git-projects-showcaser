package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// HTTPJsonFetcher is used for querying HTTP servers and fetching JSON objects
type HTTPJsonFetcher struct {
	client *http.Client
}

// NewHTTPJsonFetcher returns a new HTTPJsonFetcher object with an HTTP client initialized using the provided timeout duration
func NewHTTPJsonFetcher(timeout time.Duration) *HTTPJsonFetcher {
	return &HTTPJsonFetcher{&http.Client{Timeout: timeout}}
}

// FetchJSON fetches the JSON object from a URL into a map
func (fetcher *HTTPJsonFetcher) FetchJSON(url string) (result interface{}, err error) {
	answer, err := fetcher.client.Get(url)
	if err != nil {
		return nil, err
	}
	defer answer.Body.Close()

	if answer.StatusCode != 200 {
		return nil, fmt.Errorf("Received %v from http request", answer.Status)
	}

	var output interface{}
	err = json.NewDecoder(answer.Body).Decode(&output)

	return output, err
}
