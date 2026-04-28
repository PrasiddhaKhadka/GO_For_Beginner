package main

import (
	"fmt"
	"sync"
)

//  wait groups:

func task(id int, w *sync.WaitGroup) {
	// defer only run after the completion of func
	// defer w.Done()
	fmt.Println("Doing task: ", id)
}

func main() {

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Go(func() {
			fmt.Println("doing task ", i)
		})
		// wg.Add(1)
		// changing the value from 1 to 0 -> of same memory reference
		// go task(i, &wg)
	}

	// wg.Wait()
}
