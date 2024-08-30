package main

import "fmt"

func show_defer() {
	fmt.Printf("开始\n")
	defer fmt.Printf("1\n")
	defer fmt.Printf("2\n")
	defer fmt.Printf("3\n")
	fmt.Printf("结束\n")
}

func show_defer2() {
	fmt.Printf("开始\n")
	defer func() {
		fmt.Println("1111")
		fmt.Println("2222")
	}()
	fmt.Printf("结束\n")
}

func defer_return() int {
	a := 0
	defer func() {
		a++
		a++
	}()

	return a
}

func defer_return2() (a int) {
	defer func() {
		a++
		a++
	}()
	defer func() {
		a++
		a++
	}()
	return
}

func defer_return3() (a int) {
	b := 0
	defer func() {
		b++
		b++
	}()
	defer func() {
		b++
		b++
	}()
	return b
}

func main() {
	show_defer()
	show_defer2()
	fmt.Printf("defer_return(): %v\n", defer_return())
	fmt.Printf("defer_return2(): %v\n", defer_return2())
	fmt.Printf("defer_return3(): %v\n", defer_return3())
}
