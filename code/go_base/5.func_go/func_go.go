package main

import "fmt"

func sum_two(x, y int) int {
	return x + y
}
func sum_two2(x ...int) int {
	fmt.Printf("x - %v,type - %T\n", x, x)
	sum := 0
	for _, item := range x {
		sum += item
	}
	return sum
}

func sum_two3(x int, y ...int) int {
	sum := 0
	for _, item := range y {
		sum += item
	}
	return sum * x
}

func main() {
	fmt.Println(sum_two(10, 12))
	sum := sum_two2(1, 2, 3, 4, 5, 6) // x - [1,2,3,4,5,6] type - []int
	fmt.Print(sum)
	sum2 := sum_two3(2, 2, 3, 4, 5, 6)
	fmt.Printf("\n%v", sum2)
}
