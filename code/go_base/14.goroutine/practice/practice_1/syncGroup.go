package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func test() {

	wg.Add(1)
	for i := 0; i < 11; i++ {
		fmt.Println("test() 打印第", i, "次")
		time.Sleep(time.Microsecond * 200)
	}
	wg.Done()
}

func test2() {

	wg.Add(1)
	for i := 0; i < 11; i++ {
		fmt.Println("test2() 打印第", i, "次")
		time.Sleep(time.Microsecond * 1000)
	}
	wg.Done()
}

func main() {
	go test()
	go test2()
	for i := 0; i < 11; i++ {
		fmt.Println("main() 打印第", i, "次")
		time.Sleep(time.Millisecond * 10)
	}
	wg.Wait()
}
