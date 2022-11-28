package main

import (
	"fmt"
	"time"
)

func pass(left, right chan int) {
	left <- 1 + <-right
}

func main() {
	// The amount of go routines and channels to link together
	const n = 100000

	// Initial channel setup
	leftmost := make(chan int)
	right := leftmost
	left := leftmost

	for i := 0; i < n; i++ {
		// Creates a new channel
		right = make(chan int)

		// Links the left channel and the newly created on
		go pass(left, right)

		// makes the left channel equal to the new channel
		left = right
	}

	start := time.Now()
	// Sends a value into the right most channel
	go func(c chan int) { c <- 0 }(right)

	// Wait for it to come out on the other side
	fmt.Printf("Took %v to run %v go routines\n", time.Since(start), <-leftmost)
}
