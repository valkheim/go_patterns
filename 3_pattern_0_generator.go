package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generator(msg string) <-chan string {
	c:=make(chan string)
	go func() {
		for i := 0 ; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3) * int(time.Millisecond)))
		}	}()
	return c
}

func main() {
	c:=generator("hi")
	for i:=0 ; i < 5 ; i++ {
		// reading from the channel here is a blocking operation
		fmt.Printf("receive %q\n", <-c)
	}
	fmt.Println("leaving")
}