package main

import (
	"errors"
	"fmt"
)

func read_file(file_name string) error {
	if file_name == "main.go" {
		return nil
	} else {
		return errors.New("我们需要的是main.go文件")
	}
}

func my_fn(file_name string) {

	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("请联系管理员")
			fmt.Println(err)
		}
	}()

	if err := read_file(file_name); err != nil {
		panic(err)
	} else {
		fmt.Printf("\"读取文件\": %v\n", "读取文件")
	}
}

func main() {
	my_fn("test.go")
	fmt.Println("最后执行")
}
