package main

import (
	"fmt"
	"time"
)

func time_format() {
	time_format_str := time.Now().Format("2006-01-02 03-04-05")
	fmt.Printf("time_format_str: %v\n", time_format_str)
	time_format_str_2 := time.Now().Format("2006:01:02 03/04/05")
	fmt.Printf("time_format_str_2: %v\n", time_format_str_2)
	time_format_str_3 := time.Now().Format("2006/01/02 15/04/05")
	fmt.Printf("time_format_str_3: %v\n", time_format_str_3)

}

func main() {
	time_format()
}
