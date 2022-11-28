package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Make a channel that will signal the communication to stop
	quit := make(chan bool)

	// Service tha takes in a quit channel
	c := boring("Joe", quit)

	for i := 0; i < 5; i++ {
		fmt.Println(<-c)
	}

	fmt.Println("I'm done")

	// Send a signal into the quit channel signaling the service to stop
	quit <- true
}

// Returns a READ ONLY channel
func boring(msg string, stop <-chan bool) <-chan string {
	rand.Seed(time.Now().UnixNano())
	c := make(chan string)

	go func() {
		for i := 0; ; i++ {
			time.Sleep(time.Duration(rand.Intn(2e3)) * time.Millisecond)

			select {
			case c <- fmt.Sprintf("%s: %d", msg, i):
				// Nothing

			case <-stop:
				// Stops the service
				return
			}
		}
	}()
	return c
}
