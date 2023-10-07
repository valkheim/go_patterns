package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func worker(id int) {
	fmt.Printf("worker %d staring\n", id)
	time.Sleep(time.Duration(rand.Intn(1e3) * int(time.Millisecond)))
	fmt.Printf("worker %d ending\n", id)
}

func main(){
	var wg sync.WaitGroup
	for i := 0 ; i < 5 ; i++ {
		wg.Add(1) // increment state counter by 1
		i:=i // avoid reuse the same i value in each goroutine closure
		go func() {
			defer wg.Done() // decrement state counter by 1 (wg.Add(-1))
			worker(i) // we can use errgroup to collect worker errors
		}()
	}
	wg.Wait() // Blocks until the state counter to be null
	fmt.Println("done.")
}