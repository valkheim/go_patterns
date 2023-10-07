package main

import "fmt"

// a pipeline is a series of statges connected by channels
// where each stage is a group of goroutines running the same functions
// in each stage, the goroutines
// - receive values from upstream via inbound channels
// - work on that data
// - send values downstream via outbound channels

// Here is a generator, a producer. It is the first pipeline stage, it has no inbound channel but gives an outboud channel
func stage_0_producer(nums ...int) <-chan int {
    out := make(chan int, len(nums))
    for _, n := range nums {
        out <- n
    }
    close(out)
    return out
}

// This stage is a worker stage, it receives an inbound channel and returns an outbound channel
func stage_1_worker(in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        for n := range in {
            out <- n * n
        }
        close(out)
    }()
    return out
}

func stage_2_consumer(in <-chan int) {
    for n := range in {
        fmt.Printf("%d\n", n)
    }
}

func main() {
    stage_2_consumer(stage_1_worker(stage_0_producer(1, 2, 3, 4)))
}

/*
1
4
9
16
*/