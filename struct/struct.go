package main

import "fmt"

type Person struct {
	name    string
	age     uint8
	address string
	contact uint
	isOwner Owner
}

type Owner struct {
	isOwner  bool
	business string
	bName    string
}

func main() {

	var person Person = Person{"Ronaldo", 40, "Portugal", 9800000000, Owner{true, "Hotel", "Hotel Foot"}}
	person.contact = 9811111111
	fmt.Println(person)
	fmt.Println(person.name)
	fmt.Println(person.isOwner.bName)

	//
	fmt.Println(person.personName())
	fmt.Println(person.name)
}

// FUNCTION TIGHT WITH A STRUCT
func (person *Person) personName() string {
	person.name = "Ram Bahadur"
	return person.name
}
