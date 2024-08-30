package main

import "fmt"

// func calculator(x, y int) (int, int, int, int) {
// 	sum := x + y
// 	diff := x - y
// 	product := x * y
// 	quotient := x / y
// 	return sum, diff, product, quotient
// }

func calculator2(x, y int) (sum, diff, product, quotient int) {
	sum = x + y
	diff = x - y
	product = x * y
	quotient = x / y
	return
}

func main() {
	// a, b, c, d := calculator(2, 1)
	// fmt.Println(a)
	// fmt.Println(b)
	// fmt.Println(c)
	// fmt.Println(d)
	a, b, c, d := calculator2(2, 1)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
