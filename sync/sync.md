package main

import "fmt"

// Go runs things concurrently using goroutines.
// When multiple goroutines access the same data → you get race conditions.

// 👉 A sync.Mutex is just a lock:
// Lock() → enter critical section
// Unlock() → leave it

// ❌ Problem (race condition)

var counter int

func incremenent() {
	counter++
}

func main() {

	for i := 0; i < 1000; i++ {
		go incremenent()
	}

	fmt.Println(counter)
}
