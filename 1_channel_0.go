package main

import (
	"fmt"
	"math/rand"
	"time"
)

func boring(msg string, c chan string) {
	for i := 0 ; ; i++ {
		// sending here is a blocking operation
		c <- fmt.Sprintf("%s %d", msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3) * int(time.Millisecond)))
	}
}

func main() {
	// a channel in go provides a connection betwwen two goroutines, allowing them to communicate and synchronize
	c := make(chan string)
	go boring("hi", c)
	for i:=0 ; i < 5 ; i++ {
		// reading from the channel here is a blocking operation
		fmt.Printf("receive %q\n", <-c)
	}
	fmt.Println("leaving")
}