package main

import "fmt"

func add(a int, b int, ch chan int) {
	result := a + b
	// sending the data to the channel
	ch <- result
}

func main() {

	ch := make(chan int)

	go add(5, 5, ch)

	go add(5, 15, ch)

	// switch condition{
	// 	cond
	// }

	for {
		result := <-ch

		fmt.Println("Result:", result)

	}

}
