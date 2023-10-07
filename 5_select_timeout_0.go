package main

import (
	"fmt"
	"math/rand"
	"time"
)

func boring(msg string) <-chan string {
	c:=make(chan string)
	go func() {
		for i := 0 ; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			// If this takes more than 1 second, the consumer will timeout
			time.Sleep(time.Duration(rand.Intn(1500) * int(time.Millisecond)))
		}
	}()
	return c
}

func main() {
	c := boring("joe")
	for {
		select {
		case s:= <-c:
			fmt.Println(s)
		case <-time.After(1*time.Second): // generator for a chan time.Time
			fmt.Println("timeout!")
			return
		}
	}
}