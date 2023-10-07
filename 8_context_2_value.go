package main

import (
	"context"
	"fmt"
)

func boring(ctx context.Context) {
	fmt.Printf("boring: myKey's value is %s\n", ctx.Value("myKey"))
}

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "myKey", "myValue")
	boring(ctx)
}
