package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func putNum(intChan chan int) {
	defer wg.Done()
	for i := 0; i < 1200000; i++ {
		intChan <- i
	}
	close(intChan)
}

func primeNum(intChan chan int, primeChan chan int, exitChan chan bool) {
	defer func() {
		exitChan <- true
		wg.Done()
	}()
	for i := range intChan {
		isPrime := true
		for j := 2; j < i; j++ {
			if i%j == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			primeChan <- i
		}
	}
}

func printPrime(primeChan chan int) {
	defer wg.Done()
	for v := range primeChan {
		fmt.Println(v)
	}
}

func main() {
	start := time.Now().UnixMilli()
	intChan := make(chan int, 1000)
	primeChan := make(chan int, 1000)
	exitChan := make(chan bool, 16)
	wg.Add(1)
	go putNum(intChan)
	for i := 0; i < 16; i++ {
		wg.Add(1)
		go primeNum(intChan, primeChan, exitChan)
	}
	wg.Add(1)
	go printPrime(primeChan)
	wg.Done()

	wg.Add(1)
	go func() {
		for i := 0; i < 16; i++ {
			<-exitChan
		}
		close(primeChan)
	}()
	wg.Wait()
	end := time.Now().UnixMilli()
	fmt.Println("时间是", end-start)
	fmt.Println("执行完毕")
}
