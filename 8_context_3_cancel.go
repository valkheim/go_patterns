package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context) {
	ticker := time.NewTicker(150 * time.Millisecond)
	for {
		select {
		case <-ticker.C:
			fmt.Println("boring")
		case <-ctx.Done():
			fmt.Println("ctx.Done")
			return
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		time.Sleep(5 * 150 * time.Millisecond)
		fmt.Println("will cancel")
		cancel()
	}()

	worker(ctx)
}