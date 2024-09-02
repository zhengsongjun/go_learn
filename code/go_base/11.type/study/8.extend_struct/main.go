package main

import "fmt"

type Animal struct {
	Name string
}

func (animal Animal) Run() {
	fmt.Printf("%v在奔跑\n", animal.Name)
}

type Cat struct {
	Age    int
	Animal // 匿名字段，使得 Animal 的方法可以在 Cat 上被直接调用
}

func (cat Cat) Eat() {
	fmt.Printf("%v在吃东西\n", cat.Name)
}

type Dog struct {
	hobby []string
	*Animal
}

func (dog Dog) Play() {
	for _, value := range dog.hobby {
		fmt.Printf("%v喜欢玩: %v\n", dog.Animal.Name, value)
	}
}

func main() {
	cat := Cat{
		Animal: Animal{
			Name: "喵喵",
		},
		Age: 10,
	}
	cat.Eat()
	cat.Run() // 现在这个方法调用将正确工作
	fmt.Printf("\"_________________\": %v\n", "_________________")
	dog := Dog{
		hobby: []string{"小球", "飞盘"},
		Animal: &Animal{
			Name: "大熊",
		},
	}
	dog.Play()

}
