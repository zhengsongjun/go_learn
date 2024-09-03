package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func test(num int) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Println("这是我们的第", num, "个协程,", "打印第", i, "次")
		time.Sleep(time.Millisecond * 10)
	}
}

func main() {
	wg.Add(19)
	for i := 0; i < 19; i++ {
		go test(i)
	}
	wg.Wait()
}
