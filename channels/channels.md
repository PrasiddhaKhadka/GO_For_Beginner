package main

import (
	"fmt"
	"time"
)

func processNum(numchan chan int) {
	fmt.Println("Processing number", <-numchan)
}

func main() {

	numChan := make(chan int)

	go processNum(numChan)

	numChan <- 5

	// messageChan := make(chan string)

	// // sending data to message chan
	// messageChan <- "ping send!!!" // channels are blocking in go !

	// // receiving message
	// message := <-messageChan

	// fmt.Println("Message", message)

	time.Sleep(time.Second * 2)
}
