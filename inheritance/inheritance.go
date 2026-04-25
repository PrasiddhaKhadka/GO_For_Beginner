//  Is there inheritance in go :-> no , not !!

// can you achieve inheritance in go :-> yes , ofcourse !!!

// how :-> using composition

package main

import "fmt"

type person struct {
	name string
	age  uint8
}

func (p *person) Intro() {
	fmt.Printf("My name is %v and my age is %v\n", p.name, p.age)
}

type Employee struct {
	// COMPOSITION !!!
	*person
	employeeId int
}

func (e Employee) Intro() {
	e.person.Intro()

	fmt.Println(e.person.name)
	// inherited names
	fmt.Println(e.name)
}

func main() {

	person := &person{
		name: "Ronaldo",
		age:  40,
	}

	person.Intro()

	employe := Employee{
		person:     person,
		employeeId: 1,
	}
	employe.Intro()
}
