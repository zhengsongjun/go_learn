package main

import "fmt"

// func point_memory() {
// 	a := 10
// 	p := &a
// 	fmt.Printf("a: %v\n", a)
// 	fmt.Printf("p: %v\n", p)
// 	*p = 30
// 	fmt.Printf("a: %v\n", a)
// 	fmt.Printf("p: %v\n", p)
// }

func fn1(x int) {
	x = 10
}

func fun2(x *int) {
	*x = 30
}

func main() {
	// point_memory()
	x := 1
	fn1(x)
	fmt.Printf("x: %v\n", x)
	fun2(&x)
	fmt.Printf("x: %v\n", x)
}
