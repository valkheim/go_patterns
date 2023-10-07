package main

import (
	"fmt"
	"math/rand"
	"time"
)

// send a channel on a channel, making goroutine wait its turn

type Message struct {
	str string
	wait chan struct{} // a wait channel acts as a signaler
}

func boring(msg string) <-chan Message {
	c:=make(chan Message)
	waitForIt := make(chan struct{})
	go func() {
		for i := 0 ; ; i++ {
			c <- Message { fmt.Sprintf("%s %d", msg, i), waitForIt }
			time.Sleep(time.Duration(rand.Intn(1e3) * int(time.Millisecond)))
			<-waitForIt
		}	}()
	return c
}

func fanIn(c1, c2 <-chan Message) <-chan Message {
	// fanIn is a channel generator multiplexing c1 and c2
	c3 := make(chan Message)
	go func() { for { c3 <- <-c1}}()
	go func() { for { c3 <- <-c2}}()
	return c3
}

func main() {
	c := fanIn(boring("joe"), boring("ann"))
	for i:=0 ; i < 10 ; i++ {
		msg1 := <-c ; fmt.Println(msg1.str)
		msg2 := <-c ; fmt.Println(msg2.str)
		msg1.wait <- struct{}{}
		msg2.wait <- struct{}{}
	}
	fmt.Println("leaving")
}