package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Make a channel that will signal the communication to stop
	quit := make(chan string)

	// Service tha takes in a quit channel
	c := boring("Joe", quit)

	for i := 0; i < 5; i++ {
		fmt.Println(<-c)
	}

	// Send a signal into the quit channel signaling the service to stop
	quit <- "Please stop talking!!"

	// Wait for the service to finish cleaning up
	fmt.Printf("%v\n", <-quit)
}

func cleanUp() {
	time.Sleep(time.Duration(rand.Intn(2e3)) * time.Millisecond)
	fmt.Println("Done cleaning")
}

// Returns a READ ONLY channel
func boring(msg string, stop chan string) <-chan string {
	rand.Seed(time.Now().UnixNano())
	c := make(chan string)

	go func() {
		for i := 0; ; i++ {
			time.Sleep(time.Duration(rand.Intn(2e3)) * time.Millisecond)

			select {
			case c <- fmt.Sprintf("%s: %d", msg, i):
				// Nothing

			case <-stop:
				// Clean things up before stopping service
				cleanUp()
				stop <- "Ok, I'm finished now"
				// Stops the service
				return
			}
		}
	}()
	return c
}
