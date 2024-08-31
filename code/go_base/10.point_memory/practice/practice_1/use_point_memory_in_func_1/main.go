package main

import "fmt"

func fn1(x int) {
	x = 10
}

func fn2(x *int) {
	*x = 30
}

func main() {
	a := 1
	fn1(a)
	fmt.Printf("a: %v\n", a)
	fn2(&a)
	fmt.Printf("a: %v\n", a)
}
