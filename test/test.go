package main

import "fmt"

type Bird struct {
	name string
}

func (b *Bird) Speak() {
	fmt.Println("Bird is chirpiring. ", b)
}

type Dog struct {
	name string
}

func (d *Dog) Speak() {
	fmt.Println("Dog is barking. ", d.name)
}

type Speaking interface {
	Speak()
}

func SpeakinFunc(animal Speaking) {
	animal.Speak()
}

func main() {

	dog := Dog{
		name: "Browny!",
	}

	SpeakinFunc(&dog)

	fmt.Println("Hhhh")
}
