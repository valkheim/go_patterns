// https://go.dev/blog/context

package main

import (
	"context"
	"fmt"
)

func boring(ctx context.Context) {
	fmt.Println("boring!")
}

func main() {
	ctx := context.TODO()
	boring(ctx)
}
