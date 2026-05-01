package main

import "fmt"

type message struct {
	from     string
	to       string
	messages string
}

func receiveMessage(messageChan chan message, done chan bool) {

	defer func() { done <- true }()

	for messages := range messageChan {
		fmt.Println("Receiving message........", messages)
	}

}

func main() {

	messageChan := make(chan message, 101)
	done := make(chan bool)

	go receiveMessage(messageChan, done)

	for i := 0; i < 100; i++ {
		messageChan <- message{
			from:     "testing@gmail.com",
			to:       fmt.Sprintf("%d@gmail.com", i),
			messages: "Hello !! This is just testing email!!",
		}
	}

	close(messageChan)

	<-done

	

}
