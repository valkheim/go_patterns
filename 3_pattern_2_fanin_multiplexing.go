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
			time.Sleep(time.Duration(rand.Intn(1e3) * int(time.Millisecond)))
		}	}()
	return c
}

func fanIn(c1, c2 <-chan string) <-chan string {
	// fanIn is a channel generator multiplexing c1 and c2
	c3 := make(chan string)
	go func() { for { c3 <- <-c1}}()
	go func() { for { c3 <- <-c2}}()
	return c3
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
receive "joe 1"
receive "ann 1"
receive "joe 2"
receive "joe 3"
receive "ann 2"
receive "joe 4"
receive "joe 5"
receive "ann 3"
leaving
*/