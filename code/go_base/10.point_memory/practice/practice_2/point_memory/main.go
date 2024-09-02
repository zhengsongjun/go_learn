package main

import "fmt"

// func point_memory() {
// 	a := 10
// 	p := &a
// 	fmt.Println(a)
// 	fmt.Println(p)
// 	*p = 30
// 	fmt.Println(a)
// 	fmt.Println(p)
// }

func fn1(x int) {
	x = 10
}

func fn2(x *int) {
	*x = 30
}

func main() {
	// point_memory()
	a := 10
	fn1(a)
	fmt.Println(a)
	fn2(&a)
	fmt.Println(a)

}
