package main

import (
	"fmt"
)

//产品角色
type Car struct {
	Brand string
	Type  string
	Color string
}

func (this *Car) Drive() error {
	fmt.Printf("A %s %s %s car is running on the road!\n", this.Color, this.Type, this.Brand)
	return nil
}

//建造者角色
type Builder interface {
	PaintColor(color string) Builder
	AddBrand(brand string) Builder
	SetType(t string) Builder
	Build() Car
}

//具体的建造者
type ConcreteBuilder struct {
	ACar Car
}

func (this *ConcreteBuilder) PaintColor(cor string) Builder {
	this.ACar.Color = cor
	return this
}

func (this *ConcreteBuilder) AddBrand(bnd string) Builder {
	this.ACar.Brand = bnd
	return this
}

func (this *ConcreteBuilder) SetType(t string) Builder {
	this.ACar.Type = t
	return this
}

func (this *ConcreteBuilder) Build() Car {
	return this.ACar
}

//导演着角色
type Director struct {
	Builder Builder
}

func main() {
	dr := Director{&ConcreteBuilder{}}
	adCar := dr.Builder.SetType("SUV").AddBrand("奥迪").PaintColor("white").Build()
	adCar.Drive()

	bwCar := dr.Builder.SetType("sporting").AddBrand("宝马").PaintColor("red").Build()
	bwCar.Drive()

}
