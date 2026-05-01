package main

import "fmt"

func makeCounter() func() int {
	count := 0 // ← this variable is SHARED with the returned function

	return func() int {
		count++ // ← modifies the outer variable
		return count
	}
}

func main() {
	counter := makeCounter()

	fmt.Println(counter()) // 1
	fmt.Println(counter()) // 2
	fmt.Println(counter()) // 3

	// a NEW counter — completely separate count variable
	counter2 := makeCounter()
	fmt.Println(counter2()) // 1  ← starts fresh
}
