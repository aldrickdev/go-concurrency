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
		duration := time.Duration(rand.Intn(100)) * time.Millisecond
		time.Sleep(duration)
		return Result(fmt.Sprintf("%s result for %q in %v\n", kind, query, duration))
	}
}

// Returns the results of all of the searches
func Google(query string) []Result {
	var results []Result
	c := make(chan Result)

	// Runs all the searches in seperate go routines
	// and fans the results into 1 channel
	go func() { c <- Web(query) }()
	go func() { c <- Image(query) }()
	go func() { c <- Video(query) }()

	// Wait for the results to come from the channel
	for i := 0; i < 3; i++ {
		results = append(results, <-c)
	}

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
