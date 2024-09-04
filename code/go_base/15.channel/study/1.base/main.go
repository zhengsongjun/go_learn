package main

import "fmt"

func createChannel() {
	a := make(chan int, 3)
	for i := 0; i < 3; i++ {
		a <- i
	}
	b := <-a
	c := <-a
	d := <-a
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

}

func main() {
	createChannel()
}
