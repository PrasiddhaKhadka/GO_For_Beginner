package main

import (
	"context"
	"fmt"
)

func main() {

	// --- 1. The two root contexts ---
	// Every context in Go starts from one of these two.

	bg := context.Background()
	td := context.TODO() // placeholder when you're not sure yet which context to use

	fmt.Println("Background:", bg)
	fmt.Println("TODO:", td)

	// --- 2. What IS a context? ---
	// It's an interface that carries:
	//   - cancellation signal
	//   - deadlines / timeouts
	//   - key-value data (request-scoped values)
	// across API boundaries and goroutines.

	// --- 3. Passing context to a function ---
	doWork(bg)
}

func doWork(ctx context.Context) {
	fmt.Println("Doing work with ctx:", ctx)
}
