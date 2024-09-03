package main

import (
	"fmt"
	"runtime"
)

func main() {
	cpuNum := runtime.NumCPU()
	fmt.Printf("cpuNum: %v\n", cpuNum)
	runtime.GOMAXPROCS(cpuNum - 1)
	fmt.Println("ok")
}
