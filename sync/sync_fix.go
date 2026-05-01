package main

import (
	"fmt"
	"sync"
)

var (
	counter int
	mu      sync.Mutex
)

func increment() {
	mu.Lock()
	defer mu.Unlock()
	counter++

}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 1001; i++ {
		wg.Go(func() {
			go increment()
		})
	}
	wg.Wait()
	fmt.Println(counter)
}
