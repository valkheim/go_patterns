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
	joe:=generator("joe")
	ann:=generator("ann")
	for i:=0 ; i < 5 ; i++ {
		// bad performance if ann is ready to execute before joe, we can solve this using multiplexing
		fmt.Printf("receive %q\n", <-joe)
		fmt.Printf("receive %q\n", <-ann)
	}
	fmt.Println("leaving")
}

/*
receive "joe 0"
receive "ann 0"
receive "joe 1"
receive "ann 1"
receive "joe 2"
receive "ann 2"
receive "joe 3"
receive "ann 3"
receive "joe 4"
receive "ann 4"
leaving
*/