package main

import "fmt"

type Speak interface {
	Speak() string
}

type Dog struct{}
type Cat struct{}

func (d *Dog) Speak() string {
	return "haff"
}

func (c *Cat) Speak() string {
	return "meow"
}

// high module
type Animal struct {
	speaker Speak
}

func (animal *Animal) Make() {
	fmt.Println(animal.speaker.Speak())
}

func main() {
	dog := &Dog{}
	cat := &Cat{}

	animal1 := Animal{speaker : dog}
	animal2 := Animal{speaker : cat}
	animal1.Make()
	animal2.Make()
}


