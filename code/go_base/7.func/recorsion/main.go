package main

import "fmt"

// print n-1
func print_n_1(n int) {
	if n >= 0 {
		fmt.Printf("n: %v\n", n)
		print_n_1(n - 1)
	} else {
		return
	}
}

// calculator n - 1 的 和

func calculator_n_1(n int) int {
	if n >= 0 {
		value := n + calculator_n_1(n-1)
		return value
	} else {
		return 0
	}
}

// 阶乘计算
func factorial_n_1(n int) int {
	if n > 1 {
		value := n * factorial_n_1(n-1)
		return value
	} else {
		return 1
	}
}

func main() {
	print_n_1(10)
	fmt.Printf("calculator_n_1(10): %v\n", calculator_n_1(10))
	fmt.Printf("factorial_n_1(5): %v\n", factorial_n_1(5))
}
