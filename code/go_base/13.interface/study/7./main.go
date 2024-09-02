// 定义一个Animal接口，两个方法，SetName,GetName,让Dog和Cat实现这两个接口
package main

import "fmt"

type Animal interface {
	SetName(name string)
	GetName() string
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

func createDog() {
	var dogInterface Animal
	dog := &Dog{
		Name: "大黄",
	}
	dogInterface = dog
	fmt.Println(dogInterface.GetName())
	dogInterface.SetName("小花")
	fmt.Println(dogInterface.GetName())
}

func main() {
	createDog()
}
