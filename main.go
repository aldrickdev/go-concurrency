package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := boring("Joe")
	timeout := time.After(5 * time.Second)

	for {
		select {
		case s := <-c:
			fmt.Println(s)

		// Will timeout after 5 seconds no matter
		// how many times the loop happens
		case <-timeout:
			fmt.Println("You talk to much")
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
