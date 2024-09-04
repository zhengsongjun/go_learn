package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func setChannel(ch chan int) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		ch <- i
		time.Sleep(time.Microsecond * 500)
		fmt.Printf("\"设置成功\": %v\n", i)
	}
	close(ch)
}

func readChannel(ch chan int) {
	defer wg.Done()
	for value := range ch {
		fmt.Printf("读取成功: %v\n", value)
		time.Sleep(time.Microsecond * 10)
	}
}

func main() {
	wg.Add(2)
	myChan := make(chan int, 10)
	go setChannel(myChan)
	go readChannel(myChan)
	wg.Wait()
}
