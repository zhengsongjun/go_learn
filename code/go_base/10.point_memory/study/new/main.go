package main

import "fmt"

func point_new() {
	a := new(int)
	b := new(bool)
	fmt.Printf("a: %v\n", a)
	fmt.Printf("b: %v\n", b)
}

func main() {
	point_new()
}
