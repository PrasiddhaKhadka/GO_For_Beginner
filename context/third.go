package main

import (
	"context"
	"fmt"
)

func main() {

	// Background: the root context. Use this at the top level (main, server startup, tests)
	ctx := context.Background()

	fmt.Println(ctx)

	// TODO: a placeholder. Use when you KNOW you need context but haven't wired it up yet.
	// It signals to other developers: "this needs context, come back to this."
	ctx2 := context.TODO()
	fmt.Println(ctx2)

	// ctx, cancel := context.WithCancel(parentCtx)

	// CANCEL WHEN COMPLETED !!!!
	// defer cancel()

}
