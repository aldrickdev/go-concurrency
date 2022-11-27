package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Creates a channel of type string and passes it to the boring go routine
	c := make(chan string)
	go boring("Boring!!!", c)

	// Waits for a value to be received from the channel and prints it
	for i := 0; i < 5; i++ {
		fmt.Printf("You say: %q\n", <-c)
	}

	fmt.Println("You're boring, peace")
}

func boring(msg string, c chan string) {
	// Sends message to the channel forever
	for i := 0; ; i++ {
		c <- fmt.Sprintf("%s %d", msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}
