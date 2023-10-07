package main

import (
	"context"
	"fmt"
)

func boring(ctx context.Context) {
	fmt.Println("Doing something!")
}

func main() {
	ctx := context.Background()
	boring(ctx)
}
