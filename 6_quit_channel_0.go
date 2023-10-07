package main

import (
	"fmt"
	"math/rand"
	"time"
)

func boring(msg string, quit chan struct{}) <-chan string {
	c:=make(chan string)
	go func() {
		for {
			select {
			case c <- fmt.Sprintf("%s", msg):
				time.Sleep(time.Duration(rand.Intn(150) * int(time.Millisecond)))
			case <-quit:
				return
			}
		}
	}()
	return c
}


func main() {
	quit := make(chan struct{})
	c := boring("joe", quit)
	// Read some messages
	for i := rand.Intn(10) ; i >= 0 ; i-- {
		fmt.Printf("%s, %d more to come\n", <-c, i)
	}
	// And notify the producer to stop
	quit <- struct{}{}
}