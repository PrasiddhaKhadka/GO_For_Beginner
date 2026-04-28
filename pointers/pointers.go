package main

import (
	"errors"
	"fmt"
	"unsafe"
)

func main() {
	// Pointers()
	getAdultAgeWithPointers(nil)
}

func Pointers() {
	age := 32
	fmt.Println(unsafe.Sizeof(age))
	fmt.Printf("Age: %v \n", age)

	// WITH POINTERS
	agePointers := &age

	// prints the memory address of the value age
	fmt.Println("Age: ", agePointers)

	// prints the value from the memory address
	fmt.Println("Age: ", *agePointers)

	outputAge, err := getAdultAge(age)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Printf("Output age:  %v\n", outputAge)
	}

}

func getAdultAge(age int) (int, error) {
	if (age - 18) <= 0 {
		err := errors.New("Age must be greater than 19")
		return 0, err
	} else {
		return age - 18, nil
	}
}

// with pointers
func getAdultAgeWithPointers(age *int) (int, error) {

	// if age == nil {
	// 	return  0,errors.New("")
	// }

	// no need to return just change the value of the age
	// must deference the first one as well
	// *age = *age - 18
	if (*age - 18) <= 0 {
		err := errors.New("Age must be greater than 19")
		return 0, err
	} else {
		return *age - 18, nil
	}
}

// IMPORTANT

// takes the &pointer value from the user and update the value from the memory address
//  fmt.Scan()
