package main

import (
	"fmt"
)

func main() {
	var intArr [3]int32

	intArr[1] = 123

	fmt.Println(intArr)
	fmt.Println(intArr[0])
	fmt.Println(intArr[1:3])

	fmt.Println(&intArr[0])
	fmt.Println(&intArr[1])
	fmt.Println(&intArr[2])

	// Another way to declare
	intArrSec := [3]int8{1, 2, 3}
	fmt.Println(intArrSec)

	//
	var intSlice []int32 = []int32{4, 5, 6, 7}
	fmt.Println(intSlice)
	fmt.Printf("The length is %v with capacity %v \n", len(intSlice), cap(intSlice))
	intSlice = append(intSlice, 7)
	fmt.Println(intSlice)
	fmt.Printf("The length is %v with capacity %v \n", len(intSlice), cap(intSlice))

	// ANOTHER WAY TO MAKE
	var intSlice3 []int32 = make([]int32, 3, 8)
	fmt.Println(intSlice3)

	// MAP
	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

	myMap2 := map[string]string{"name": "Ronaldo", "age": "40", "country": "Portugal"}
	fmt.Println(myMap2["name"])

	delete(myMap2, "country")

	var age, ok = myMap2["age"]
	if ok {
		fmt.Printf("The age is %v\n", age)
	} else {
		fmt.Println("The age not found")
	}

}
