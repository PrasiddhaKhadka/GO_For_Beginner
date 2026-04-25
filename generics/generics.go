package main

import "fmt"

//  GENERIC FUN

// 🤞 DON'T USE any !!!!

func displayItem[T int | string | bool](items []T) {

	for _, val := range items {
		fmt.Println(val)
	}

}

//  FOR TYPE

type stack[T int | string | bool] struct {
	elements []T
}

func main() {

	// items := []string{"1", "2", "3", "4", "5"}
	itemsInt := []int{1, 2, 3, 4, 5}

	// stack := stack[string]{
	// 	elements: []string{"golang"},
	// }

	displayItem(itemsInt)

}
