// Part A — A Regular Function vs A Method

package main

import "fmt"

// regular function — standalone, not attached to anything
func greet(name string) {
	fmt.Println("Hello", name, ":::simple func:::")
}

// a struct
type Person struct {
	name string
}

// method — attached to Person via a receiver (p Person)
func (p Person) reName(newName string) {
	p.name = newName // ← changes a copy, original untouched
	fmt.Println("Hello", p.name)
	fmt.Println("Name changed temporarly!!")

}

func (p *Person) renamePermanently(newName string) {
	p.name = newName
	fmt.Println("Name changed permanently!!")
}

func main() {
	// regular function — you pass data IN
	greet("Ram")

	// method — the data is already INSIDE the struct
	p := Person{name: "Ram"}
	p.reName("New Ram")
	fmt.Println(p.name)
	fmt.Println("After temp")

	//
	p.renamePermanently("James")
	fmt.Println(p.name)

}
