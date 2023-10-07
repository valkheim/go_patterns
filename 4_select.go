package main

import (
	"fmt"
	"math/rand"
	"time"
)

// a control structure unique to concurrency
// the reason channels and goroutines are built into the language rather than just a library

// it's like a switch but each case is a communication
// all channels are evaluated
// selection blocks until one communication can proceed
// if multiple can proceed, select chooses pseudo-randomly
// a default case, if present, executes immediately if no channel is ready (we can make channels non blocking that way)

// We can rewrite the fanIn of the multiplexer example using select
func fanIn(c1, c2 <-chan string) <-chan string {
	c3 := make(chan string)
	go func() {
		for {
			select {
			case s:= <-c1:
				c3 <- s
			case s:= <-c2:
				c3 <- s
			}
		}
	}()
	return c3
}

func boring(msg string) <-chan string {
	c:=make(chan string)
	go func() {
		for i := 0 ; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3) * int(time.Millisecond)))
		}	}()
	return c
}

func main() {
	c := fanIn(boring("joe"), boring("ann"))
	for i:=0 ; i < 10 ; i++ {
		fmt.Printf("receive %q\n", <-c)
	}
	fmt.Println("leaving")
}

/*
receive "joe 0"
receive "ann 0"
receive "ann 1"
receive "ann 2"
receive "joe 1"
receive "ann 3"
receive "ann 4"
receive "ann 5"
receive "joe 2"
receive "ann 6"
leaving
*/