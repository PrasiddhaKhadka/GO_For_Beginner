package main

import (
	"fmt"
)

type Speaker interface {
	Speak(message string) string
}

type Dog struct {
	name string
	age  int
}

type Cat struct {
	name  string
	color string
}

type Bird struct {
	name  string
	speed int
}

func (d Dog) Speak(message string) string {
	return d.name + " says: " + message
}

func (c Cat) Speak(message string) string {
	return c.name + " says: " + message
}

func (b Bird) Speak(message string) string {
	return b.name + " says: " + message
}

func main() {

	animals := []Speaker{
		Dog{name: "Rex", age: 5},
		Cat{name: "Mimi", color: "orange"},
		Dog{name: "Bruno", age: 3},
		Bird{name: "Chirpir", speed: 50},
	}

	for _, animal := range animals {
		fmt.Println(animal.Speak("hello!"))

		// if dog, ok := animal.(Dog); ok {
		// 	fmt.Println("  → This is a Dog, age:", dog.age)
		// }

		// if cat, ok := animal.(Cat); ok {
		// 	fmt.Println("  → This is a Cat, color:", cat.color)
		// }

		// if bird, ok := animal.(Bird); ok {
		// 	fmt.Println("  → This is a Bird, color:", bird.speed)
		// }

		switch v := animal.(type) {
		case Dog:
			fmt.Println(v.age)
		case Cat:
			fmt.Println(v.color)
		case Bird:
			fmt.Println(v.speed)
		}
	}

}
