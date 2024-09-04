package main

import "fmt"

func forRange() {
	a := make(chan int, 3)
	for i := 0; i < 3; i++ {
		a <- i
	}
	close(a)
	for value := range a {
		fmt.Println(value)
	}
}

func main() {
	forRange()
}
