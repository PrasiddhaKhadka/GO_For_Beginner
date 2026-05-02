package main

//"I don't care what TYPE you are — I only care what you CAN DO"

import "fmt"

// "anything that has Speak() string — is a Speaker"
type Speaker interface {
Speak() string
}

type Dog struct {
Name string
}

type Cat struct {
Name string
}

type Bird struct {
Name string
}

func (d Dog) Speak() string {
return d.Name + " THE DOG: " + "Woof!"
}

func (c Cat) Speak() string {
return c.Name + " THE CAT: " + "Meow!"
}

func (b Bird) Speak() string {
return b.Name + " says: Tweet!"
}

// --- ONE function for ALL animals ---
func MakeItSpeak(animal Speaker) {
fmt.Println(animal.Speak())
}

func main() {

    // s holds TWO things internally:
    // 1. the actual value  → Dog{Name: "Rex"}
    // 2. the actual type   → Dog
    var s Speaker = Dog{Name: "Rex"}

    // dog := Dog{Name: "Rex"}
    cat := Cat{Name: "Mimi"}

    bird := Bird{Name: "Tweety"}

    // fmt.Println(dog.Speak())
    fmt.Println(cat.Speak())

    MakeItSpeak(cat)
    MakeItSpeak(bird)

    MakeItSpeak(s)

    if dog, ok := s.(Cat); ok {
    	fmt.Println("  → This is a Dog,", dog.Name)
    }

    // MakeItSpeak(cat)

}
