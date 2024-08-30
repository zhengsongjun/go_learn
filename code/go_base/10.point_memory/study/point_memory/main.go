package main

import "fmt"

func point_memory() {
	a := 10
	p := &a
	fmt.Printf("a: %v\n", a)
	fmt.Printf("p: %v\n", p)

	*p = 30
	fmt.Printf("a: %v\n", a)

}

func main() {
	point_memory()
}
