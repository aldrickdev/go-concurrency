package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	go boring("Boring!!!")
	// The program will finish immediately as once the
	//   go routine is launched, main finishes
}

// Non-concurrent function
func boring(msg string) {
	for i := 0; ; i++ {
		fmt.Println(msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}
