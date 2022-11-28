package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Define the different searches
var (
	Web1 = fakeSearch("web1")
	Web2 = fakeSearch("web2")

	Image1 = fakeSearch("image1")
	Image2 = fakeSearch("image2")

	Video1 = fakeSearch("video1")
	Video2 = fakeSearch("video2")
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

	// Runs all the replicated searches in seperate go routines
	// and fans the results into 1 channel
	go func() { c <- FirstResult(query, Web1, Web2) }()
	go func() { c <- FirstResult(query, Image1, Image2) }()
	go func() { c <- FirstResult(query, Video1, Video2) }()

	// Sets up the timeout
	timeout := time.After(80 * time.Millisecond)

	// Wait for the 3 results to come back but if the total elapsed time
	// is over the set timeout quit the search and return anyways
	for i := 0; i < 3; i++ {
		select {
		case r := <-c:
			results = append(results, r)

		case <-timeout:
			fmt.Println("Timed Out")
			return results
		}
	}

	return results
}

// Returns the first value that the replicas find
func FirstResult(query string, replicas ...Search) Result {
	c := make(chan Result)

	// Creates an inline function that will run the search query
	// and puts the result in the channel
	searchReplica := func(i int) { c <- replicas[i](query) }

	// Runs each of the replicas in a go routine
	for i := range replicas {
		go searchReplica(i)
	}

	// returns the fist value that comes back
	return <-c
}

func main() {
	rand.Seed(time.Now().UnixNano())

	start := time.Now()
	results := Google("golang")
	elapsed := time.Since(start)

	fmt.Println(results)
	fmt.Println(elapsed)
}
