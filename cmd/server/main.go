package main

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

func main() {

	fmt.Println("Hello World")

	// DECLEARING THE VAR NAME
	var intNum int
	fmt.Println(intNum)

	var floatNum float32
	fmt.Println(floatNum)

	var floarNum32 float32 = 10.1
	var intNum32 int32 = 2

	var result float32 = floarNum32 + float32(intNum32)
	fmt.Println(result)

	var intNum1 int = 3
	var intNum2 int = 2
	fmt.Println(intNum1 / intNum2)
	fmt.Println(intNum1 % intNum2)

	var myString string = `Hello World`
	fmt.Println(myString)
	fmt.Println(len(myString))

	fmt.Println((len("A")))
	// TAKING THE CHAR OUT OF LETTER LIKE GAMMA
	fmt.Println((len("γ")), "gamma!")

	// changing the special char
	fmt.Println(utf8.RuneCountInString("γ"))

	var myRune rune = 'a'
	fmt.Println(myRune)

	var myBoolean bool = false
	fmt.Println(myBoolean)

	// var myText = "Hello World"      : valid also
	// myText := "Hello World"         : valid no need to use var

	printMe("Hello World")

	var output, remainder, err = intDivision(10, 0)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else if remainder == 0 {
		fmt.Printf("The result of the integer division is %v ", output)
	} else {
		fmt.Printf("The result of the integer division is %v with remainder %v\n", output, remainder)
	}

}

func printMe(printValue string) {
	fmt.Println(printValue)
}

func intDivision(numerator int, denominator int) (int, int, error) {
	// error return type is nil by default
	var error error
	if denominator == 0 {
		err := errors.New("Cannot divide by zero")
		return 0, 0, err
	}
	var result int = numerator / denominator
	var remainder int = numerator % denominator
	return result, remainder, error
}
