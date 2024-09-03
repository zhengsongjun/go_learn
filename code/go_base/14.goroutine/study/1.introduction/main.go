package main

import (
	"fmt"
	"time"
)

func test() {
	for i := 0; i < 10; i++ {
		fmt.Println("test() 执行 - ", i)
		time.Sleep(time.Millisecond * 300)
	}
}

func main() {
	go test()
	for i := 0; i < 10; i++ {
		fmt.Println("main() 执行 - ", i)
		time.Sleep(time.Millisecond * 300)
	}
}
