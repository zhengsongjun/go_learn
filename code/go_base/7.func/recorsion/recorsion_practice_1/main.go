package main

import "fmt"

// print n-1
// func print_n_1(n int) {
// 	if n > 0 {
// 		fmt.Printf("n: %v\n", n)
// 		print_n_1(n - 1)
// 	} else {
// 		fmt.Printf("n: 0\n")
// 	}
// }

// calculator n - 1 的 和

// func calcutator(n int) int {
// 	value := 0
// 	if n > 0 {
// 		value = n + calcutator(n-1)
// 	} else {
// 		return 0
// 	}
// 	return value
// }

func factorial_n(n int) int {
	value := 0
	if n > 1 {
		value = n * factorial_n(n-1)
	} else {
		return 1
	}
	return value
}

// 阶乘计算

func main() {
	// print_n_1(100)
	// fmt.Printf("calcutator(5): %v\n", calcutator(5))
	// fmt.Printf("calcutator(100): %v\n", calcutator(100))
	fmt.Printf("factorial_n(4): %v\n", factorial_n(4))

}
