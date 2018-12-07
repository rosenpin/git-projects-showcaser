package utils

import (
	"fmt"
	"testing"
	"time"
)

func TestHTTPJsonFetcher(t *testing.T) {
	expected := map[string]interface{}{"id": 1, "title": "delectus aut autem", "completed": false, "userId": 1}

	fetcher := NewHTTPJsonFetcher(10 * time.Second)
	result, err := fetcher.FetchJSON("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		t.Error("Failed fetching json from fake website, err: ", err)
	}

	fmt.Printf("%T\n", result)
	parsedResult := result.(map[string]interface{})

	for key, value := range expected {
		if fmt.Sprint(parsedResult[key]) != fmt.Sprint(value) {
			t.Errorf("Unexpeted result, expected %v:%v, received %v:%v", key, value, key, fmt.Sprint(parsedResult[key]))
		}
	}
}
