package main

import (
	"fmt"
)

type Action interface {
	Move(int) int
}

type Animal struct {
	Name string
}

type Bird struct {
	Animal
}

func (this *Bird) Move(num int) int {
	fmt.Printf("I am a bird, I flyed %d meters.\n", num)
	return num
}

type Fish struct {
	Animal
}

func (this *Fish) Move(num int) int {
	fmt.Printf("I am a fish, I swimmed %d meters.\n", num)
	return num
}

type Dog struct {
	Animal
}

func (this *Dog) Move(num int) int {
	fmt.Printf("I am a dog, I flying %d meters.\n", num)
	return num
}

type AnimalFactory struct {
}

func NewAnimalFactory() *AnimalFactory {
	return &AnimalFactory{}
}

func (this *AnimalFactory) CreateAnimal(name string) Action {
	switch name {
	case "bird":
		return &Bird{}
	case "fish":
		return &Fish{}
	case "dog":
		return &Dog{}
	default:
		panic("error animal type")
		return nil
	}
}

func main() {
	bird := NewAnimalFactory().CreateAnimal("bird")
	bird.Move(100)

	fish := NewAnimalFactory().CreateAnimal("fish")
	fish.Move(200)

	dog := NewAnimalFactory().CreateAnimal("dog")
	dog.Move(300)
}
