package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := fanIn(boring("Joe"), boring("Ann"))

	// Pull from the channels 5 times
	for i := 0; i < 5; i++ {
		// Gets the Message from the channel and prints the str
		msg1 := <-c
		fmt.Println(msg1.str)

		msg2 := <-c
		fmt.Println(msg2.str)

		// Now that we got both messages, we can tell them to continue
		msg1.wait <- true
		msg2.wait <- true
	}

	fmt.Println("You're both boring, peace")
}

// Returns a READ ONLY channel
func boring(msg string) <-chan Message {
	// Creates the channels
	c := make(chan Message)
	waitForIt := make(chan bool)

	// Launches a go routine that will place values into the channel
	go func() {
		// Sends message to the channel forever
		for i := 0; ; i++ {
			// Sends the message and a wait channeled used to wait for a signal to continue
			c <- Message{fmt.Sprintf("%s: %d", msg, i), waitForIt}
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
			<-waitForIt
		}
	}()
	return c
}

type Message struct {
	str  string
	wait chan bool
}

func fanIn(input1, input2 <-chan Message) <-chan Message {
	c := make(chan Message)

	go func() {
		for {
			c <- <-input1
		}
	}()
	go func() {
		for {
			c <- <-input2
		}
	}()

	return c
}
