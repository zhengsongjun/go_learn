package main

import "fmt"

type calc func(int, int) int
type kelihua func(int) int

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

func getCalc(x string) calc {
	switch x {
	case "+":
		return add
	case "-":
		return sub
	default:
		return add
	}
}

func kelihua_func(x int) kelihua {
	return func(y int) int {
		return x + y
	}
}

func main() {
	calc_value := getCalc("+")
	fmt.Printf("calc_value(1, 2): %v\n", calc_value(1, 2))
	get_calc := getCalc("-")
	fmt.Printf("get_calc: %v\n", get_calc(3, 1))
	kelihua_value := kelihua_func(100)
	fmt.Printf("kelihua_value(200): %v\n", kelihua_value(200))
}
