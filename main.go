package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := boring("Joe")

	for {
		select {
		case s := <-c:
			fmt.Println(s)

		// Creates a channel that will return a value when the time has elapsed
		// Note that this timer will restart after every loop
		case <-time.After(1 * time.Second):
			fmt.Println("You're too slow")
			return
		}
	}
}

// Returns a READ ONLY channel
func boring(msg string) <-chan string {
	rand.Seed(time.Now().UnixNano())
	c := make(chan string)

	go func() {
		for {
			t := time.Duration(rand.Intn(2e3))
			c <- fmt.Sprintf("%s\t%v", msg, t)
			time.Sleep(t * time.Millisecond)
		}
	}()
	return c
}
