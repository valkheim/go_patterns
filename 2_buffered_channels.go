package main

import "fmt"

func main() {
	// go channels can be created with a buffer
	// buffering removes synchronization,
	// they have to property that they don't synchronize when you send becuase you can just drop a value in the buffer and keep going

	msg := make(chan string, 2)
	msg <- "push first msg into buffer"
	msg <- "push second msg into buffer without a received having to read the first msg"

	fmt.Println(<-msg)
	fmt.Println(<-msg)
}