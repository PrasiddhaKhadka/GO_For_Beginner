package main

import (
	"fmt"
	"math/rand"
	"sync"
)

// 1. Important concept ( CHANNELS ARE BLOCKING !!!) :-> sender and receiver both must be ready to communicate else deadlock occurs
// func main() {

// 	msgChan := make(chan string)

// 	// Receiving data to a channel

// 	msgChan <- "Ping"

// 	// sending data

// 	rec := <-msgChan

// 	fmt.Println(rec)
// }

// 2.  WAY TO FIX THE DEADLOCK

func processNum(ch chan int) {
	for num := range ch {
		fmt.Println("Processing the number", num)
	}
}

// main go routine()
func main() {

	var wg sync.WaitGroup

	numchan := make(chan int)

	// another go routine
	wg.Go(func() {
		processNum(numchan)
	})

	// numchan <- 5

	for {
		numchan <- rand.Intn(1000)
	}

	// wg.Wait()

}

// Goroutine #1 → main
// Goroutine #2 → processNum
