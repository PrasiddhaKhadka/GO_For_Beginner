//

package main

import "fmt"

type Speaker interface {
	Speak(message string) string
}

type Dog struct {
	name string
}

func (d Dog) Speak(message string) string {
	return d.name + " " + message
}

type Cat struct {
	name string
}

func (c Cat) Speak(message string) string {
	return c.name + " " + message
}

type Bird struct {
	name string
}

func (b Bird) Speak(message string) string {
	return b.name + " " + message
}

func MakeItSpeak(animal Speaker, message string) {
	fmt.Println(animal.Speak(message))

}

func main() {
	animals := []Speaker{
		Dog{name: "Rex"},
		Cat{name: "Mimi"},
		Bird{name: "Tweety"},
		Dog{name: "Bruno"},
	}

	for _, data := range animals {
		MakeItSpeak(data, "Says hello!!!")
	}

	fmt.Println(animals[0].Speak("Says Dance"))
}
