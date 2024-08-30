package main

import "fmt"

type calc func(int, int) int
type calc2 func(int) int

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

func computed(x int, y int, method calc) int {
	return method(x, y)
}

func getCalc(clac string) calc {
	switch clac {
	case "+":
		return add
	case "-":
		return sub
	default:
		return add
	}
}

func kelihua(x int) calc2 {
	return func(y int) int {
		return x + y
	}
}

func main() {
	fmt.Printf("computed(1, 2, add): %v\n", computed(1, 2, add))
	calc_func := getCalc("+")
	fmt.Printf("calc_func(2, 1): %v\n", calc_func(2, 1))
	calc_func2 := getCalc("-")
	fmt.Printf("calc_func2: %v\n", calc_func2(2, 1))

	func_kelihua := kelihua(1)
	fmt.Printf("func_kelihua(2): %v\n", func_kelihua(2))
}
