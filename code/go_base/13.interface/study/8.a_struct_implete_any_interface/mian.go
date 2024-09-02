package main

import "fmt"

type AnimalSetName interface {
	SetName(name string)
}

type AnimalGetName interface {
	GetName() string
}

type Dog struct {
	Name string
}

func (a *Dog) SetName(name string) {
	a.Name = name
}

func (a Dog) GetName() string {
	return a.Name
}

func main() {

	var aGetName AnimalGetName
	var aSetName AnimalSetName
	dog := &Dog{
		Name: "大黄",
	}

	aGetName = dog
	aSetName = dog
	fmt.Printf("aGetName.GetName(): %v\n", aGetName.GetName())
	aSetName.SetName("大熊")
	fmt.Printf("aGetName.GetName(): %v\n", aGetName.GetName())

}
