package main

import "fmt"

type myInt int

func (my_int myInt) PrintInfo() {
	fmt.Printf("我是自定义类型，我的类型是myInt,我的值是%v", my_int)
}

func main() {
	var a myInt = 1
	a.PrintInfo()
}
