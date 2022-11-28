package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Define the different searches
var (
	Web   = fakeSearch("web")
	Image = fakeSearch("image")
	Video = fakeSearch("video")
)

// Define some types
type Result string
type Search func(query string) Result

// Create a closure for the search functions
func fakeSearch(kind string) Search {
	return func(query string) Result {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return Result(fmt.Sprintf("%s result for %q\n", kind, query))
	}
}

// Returns the results of all of the searches
func Google(query string) []Result {
	var results []Result

	results = append(results, Web(query))
	results = append(results, Image(query))
	results = append(results, Video(query))

	return results
}

func main() {
	rand.Seed(time.Now().UnixNano())

	start := time.Now()
	results := Google("golang")

	elapsed := time.Since(start)
	fmt.Println(results)
	fmt.Println(elapsed)
}
