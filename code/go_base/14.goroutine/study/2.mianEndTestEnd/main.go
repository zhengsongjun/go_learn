package main

import (
	"fmt"
	"time"
)

func test() {
	for i := 0; i < 10; i++ {
		fmt.Println("test 打印 - ", i)
		time.Sleep(time.Millisecond * 1000)
	}
}

func main() {
	go test()
	for i := 0; i < 10; i++ {
		fmt.Println("main 打印 - ", i)
		time.Sleep(time.Millisecond * 200)
	}
	// 主进程结束了，主进程其他协程不管有没有执行完，都会跟着结束
}
