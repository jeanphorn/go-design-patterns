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

type AnimalFactory interface {
	CreateAnimal() Action
}

type BirdFactory struct {
}

func (this *BirdFactory) CreateAnimal() Action {
	return &Bird{}
}

type FishFactory struct {
}

func (this *FishFactory) CreateAnimal() Action {
	return &Fish{}
}

func main() {
	bFactory := &BirdFactory{}
	bird := bFactory.CreateAnimal()
	bird.Move(100)

	fFactory := &FishFactory{}
	fish := fFactory.CreateAnimal()
	fish.Move(200)

}
