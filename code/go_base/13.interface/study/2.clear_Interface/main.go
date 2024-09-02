package main

import "fmt"

type A interface {
}

func clearInterface() {
	var a A
	s := "sss"
	a = s
	fmt.Printf("a的值:%v,a的类型:%T\n", a, a)

	num := 20
	a = num
	fmt.Printf("a的值:%v,a的类型:%T\n", a, a)
}

func anyType() {
	var a interface{}

	a = 10
	a = false
	a = "sss"
	a = []int{1, 2, 3, 4, 5}

	fmt.Printf("a:%v\n", a)
}

func main() {
	clearInterface()
	anyType()
}
