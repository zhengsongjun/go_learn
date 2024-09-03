package main

import (
	"fmt"
	"sync"
	"time"
)

var w sync.WaitGroup

func test1() {
	w.Add(1)
	for i := 0; i < 10; i++ {
		fmt.Println("test1 打印 - ", i)
		time.Sleep(time.Millisecond * 300)
	}
	w.Done()
}

func test2() {
	w.Add(1)
	for i := 0; i < 10; i++ {
		fmt.Println("test2 打印 -", i)
		time.Sleep(time.Millisecond * 500)
	}
	w.Done()
}

func main() {
	go test1()
	go test2()
	for i := 0; i < 10; i++ {
		fmt.Println("main 打印 -", i)
		time.Sleep(time.Millisecond * 10)
	}
	w.Wait()
	fmt.Println("主线程执行完毕")
}
