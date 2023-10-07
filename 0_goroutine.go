package main

import (
	"fmt"
	"math/rand"
	"time"
)

func boring(msg string) {
	for i := 0 ; ; i++ {
		fmt.Println(msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3) * int(time.Millisecond)))
	}
}

func main() {
	// In go, concuerrency is the composition of independently executing goroutines

	// A goroutine is an independently executing function, launched by a go statement
	// It has its own call stack
	// It's very cheap, you can have 100k+
	// It's not a thread
	// Goroutines are multiplexed dynamically onto threads as needed to keep all goroutines running
	go boring("hi")

	// When main exists, the program exits so let's delay a bit
	time.Sleep(2*time.Second)
}