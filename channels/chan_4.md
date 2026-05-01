// UNBUFFERED !
package main

import (
	"fmt"
	"sync"
)

func sum(numchan chan int, a int, b int) {
	result := a + b
	// sending the result to the channel
	numchan <- result
}

func main() {

	var wg sync.WaitGroup

	numChan := make(chan int)

	wg.Go(func() {
		sum(numChan, 5, 5)
	})

	// RECEIVING THE DATA
	result := <-numChan

	fmt.Println(result)

	wg.Wait()

}
