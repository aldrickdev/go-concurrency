package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Launch 2 services with different names
	joe := boring("Joe")
	ann := boring("Ann")

	// Pull from the channels 5 times
	for i := 0; i < 5; i++ {
		fmt.Println(<-joe)
		fmt.Println(<-ann)
	}

	fmt.Println("You're both boring, peace")
}

// Returns a READ ONLY channel
func boring(msg string) <-chan string {
	// Creates the channel
	c := make(chan string)

	// Launches a go routine that will place values into the channel
	go func() {
		// Sends message to the channel forever
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}
