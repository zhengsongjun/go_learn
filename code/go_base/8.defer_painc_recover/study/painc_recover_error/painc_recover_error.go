package main

import "fmt"

func recover_error(x, y int) {
	defer func() {
		error := recover()
		if error != nil {
			fmt.Printf("error: %v\n", error)
		}
	}()
	fmt.Println(x / y)
}

func main() {
	recover_error(10, 0)
	fmt.Printf("\"继续执行\": %v\n", "继续执行")
	recover_error(2, 1)
}
