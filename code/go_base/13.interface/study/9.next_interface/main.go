package main

import "fmt"

type AnimalSetName interface {
	SetName(name string)
}

type AnimalGetName interface {
	GetName() string
}

type Animal interface {
	AnimalSetName
	AnimalGetName
}

type Dog struct {
	Name string
}

func (d *Dog) SetName(name string) {
	d.Name = name
}

func (d Dog) GetName() string {
	return d.Name
}

func main() {
	var a Animal
	dog := Dog{
		Name: "dog",
	}

	a = &dog
	fmt.Printf("dog.GetName(): %v\n", a.GetName())
	a.SetName("小 dog")
	fmt.Printf("dog.GetName(): %v\n", a.GetName())
}
