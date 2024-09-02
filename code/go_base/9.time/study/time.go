package main

import (
	"fmt"
	"time"
)

func time_base() {
	time_obj := time.Now()
	fmt.Printf("time_obj: %v\n", time_obj) //2024-08-30 16:57:38.172531 +0800 CST m=+0.000149926

	year := time_obj.Year()
	month := time_obj.Month()
	day := time_obj.Day()
	hour := time_obj.Hour()
	minute := time_obj.Minute()
	second := time_obj.Second()

	fmt.Printf("%d-%d-%d %d:%d:%d \n", year, month, day, hour, minute, second)
	fmt.Printf("%02d-%02d-%02d %02d:%02d:%02d", year, month, day, hour, minute, second)
}

func main() {
	time_base
}
