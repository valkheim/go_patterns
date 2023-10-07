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
			time.Sleep(time.Duration(rand.Intn(150) * int(time.Millisecond)))
		}
	}()
	return c
}

func main() {
	c := boring("joe")
	overallTimeout := time.After(1*time.Second)  // generator for a chan time.Time
	for {
		select {
		case s:= <-c:
			fmt.Println(s)
		case <- overallTimeout: // this will fire 1 second after the timeout creation
			fmt.Println("timeout!")
			return
		}
	}
}