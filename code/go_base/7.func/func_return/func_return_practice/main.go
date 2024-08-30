package main

import "fmt"

func func_return_no_name(x, y int) (int, int, int, int) {
	sum := x + y
	diff := x - y
	product := x * y
	quotient := x / y
	return sum, diff, product, quotient
}

func func_return_name(x, y int) (sum, diff, product, quotient int) {
	sum = x + y
	diff = x - y
	product = x * y
	quotient = x / y
	return
}

func main() {
	sum, diff, product, quotient := func_return_no_name(2, 1)
	fmt.Println(sum, diff, product, quotient)
	a, b, c, d := func_return_name(2, 1)
	fmt.Println(a, b, c, d)
}
