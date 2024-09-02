package main

import "fmt"

// 自定义类型
type myInt int

// 类型别名
type intOtherName = int

func main() {
	var a myInt = 10        // main.myInt
	var b intOtherName = 10 // int
	fmt.Printf("a: %T\n", a)
	fmt.Printf("b: %T\n", b)
}
