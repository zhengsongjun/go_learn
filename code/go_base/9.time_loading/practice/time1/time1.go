package main

import (
	"fmt"
	"time"
)

func time_base() {
	time_obj := time.Now()
	year := time_obj.Year()
	month := time_obj.Month()
	day := time_obj.Day()
	hour := time_obj.Hour()
	minute := time_obj.Minute()
	second := time_obj.Second()
	fmt.Printf("%d-%d-%d %d:%d:%d\n", year, month, day, hour, minute, second)
	fmt.Printf("%02d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
}

func main() {
	time_base()
}
