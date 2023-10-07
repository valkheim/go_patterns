package main

import (
	"fmt"
	"time"
)

func demo_1() {
	ch := make(chan int, 1)
    ch <- 2
    val, ok := <-ch
    fmt.Printf("Val: %d OK: %t\n", val, ok) // channel is still open (ok)

    close(ch)
    val, ok = <-ch
    fmt.Printf("Val: %d OK: %t\n", val, ok) //  channel is closed (!ok)
}


func sum(ch chan int) {
	sum := 0
	for val := range ch { // consume until ch is closed
		sum += val
	}
	fmt.Printf("Sum: %d\n", sum)
}

func demo_2() {
	ch := make(chan int, 3)
	ch <- 2
	ch <- 2
	ch <- 2
	close(ch) // required so the range in sum can be notified by the close
	sum(ch)
	time.Sleep(time.Second * 1)

}

func main() {

	demo_1()
	fmt.Println("##############")
	demo_2()

}