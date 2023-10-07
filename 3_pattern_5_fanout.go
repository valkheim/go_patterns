// https://go.dev/blog/pipelines
// stages close their outbound channels when all the send operations are done.
// stages keep receiving values from inbound channels until those channels are closed or the senders are unblocked.

package main

import (
	"fmt"
	"sync"
)

func generator(done<-chan struct{}, nums ...int) <-chan int {
	// create outbound channel
    out := make(chan int)
    go func() {
        for _, n := range nums {
			// send data to outbound channel
			select {
			case out <- n:
			case <-done:
			}
		}
		// When everything is sent, close the stage 0 outbound channel
        close(out)
    }()
    return out
}


func worker(done <-chan struct{}, in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
		defer close(out)
        for n := range in {
			select {
			case out <- n * n:
			case <-done:
				return
			}
        }
    }()
    return out
}

// Merge multiple cs channels into one outbound channel
func fanOut(done <-chan struct{}, cs ...<-chan int) <-chan int {
	// create and outbound channel
    out := make(chan int)

	// We'll have to wait until data from every inbound 'cs' channels is sent to the 'out' outbound channel
	var wg sync.WaitGroup
	wg.Add(len(cs))
	
	for _, c := range cs {
		// Start an output goroutine for each input channel in cs.
		// Copy values from c to out until c is closed
        go func(c <-chan int) {
			defer wg.Done()
			for n := range c {
				select {
					case out <- n:
					case <-done:
						return
				}
			}
		}(c)
    }

    // Start a goroutine to close out once all the output goroutines are done.
    go func() {
        wg.Wait()
        close(out)
    }()

	return out
}

func main() {
	done := make(chan struct{})

	// create inbound goroutine
	in := generator(done, 1,2,3,4)

	// distribute the work among multiple goroutines
	c1 := worker(done, in)
	c2 := worker(done, in)

	// collect c1,c2 outputs in a single outbound channel
	for n := range fanOut(done, c1, c2) {
		fmt.Printf("%d\n", n)
	}

	close(done)

	fmt.Println("done.")
}

/*
1
4
9
16
done.
*/