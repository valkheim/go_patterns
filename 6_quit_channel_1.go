package main

import (
	"fmt"
	"math/rand"
	"time"
)

func boring(msg string, quit chan string) <-chan string {
	c:=make(chan string)
	go func() {
		for {
			select {
			case c <- fmt.Sprintf("%s", msg):
				time.Sleep(time.Duration(rand.Intn(150) * int(time.Millisecond)))
			case <-quit:
				// do stuff
				quit <- "bye consumer"
				return
			}
		}
	}()
	return c
}


func main() {
	quit := make(chan string)
	c := boring("producer", quit)
	// Read some messages
	for i := rand.Intn(10) ; i >= 0 ; i-- {
		fmt.Printf("%s, %d more to come\n", <-c, i)
	}
	// Notify the producer to stop, wait for its ack // round trip communication
	quit <- "bye producer"
	fmt.Printf("producer says %s\n", <-quit)
}