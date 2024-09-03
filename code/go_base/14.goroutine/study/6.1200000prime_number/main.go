package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func test(n int) {
	defer wg.Done()
	for i := (n - 1) * 30000; i < n*30000; i++ {
		if i > 1 {
			isPrime := true
			for j := 2; j < i; j++ {
				if i%j == 0 {
					isPrime = false
					break
				}
			}
			if isPrime {

			}
		}
	}

}

func test2() {
	start := time.Now().UnixMilli() // 使用 UnixMilli() 获取毫秒级时间戳
	for i := 0; i < 120000; i++ {
		if i > 1 {
			isPrime := true
			for j := 2; j < i; j++ { // 修正为 j++
				if i%j == 0 {
					isPrime = false
					break
				}
			}
			if isPrime {
				// 这里可以放置其他处理代码
			}
		}
	}
	end := time.Now().UnixMilli() // 使用 UnixMilli() 获取毫秒级时间戳
	fmt.Println(end-start, "毫秒")
}

func main() {
	test2()
	start := time.Now().UnixMilli()
	for i := 0; i <= 4; i++ {
		wg.Add(1)
		go test(i)
	}
	wg.Wait()
	end := time.Now().UnixMilli()
	fmt.Println(end-start, "毫秒")

}
