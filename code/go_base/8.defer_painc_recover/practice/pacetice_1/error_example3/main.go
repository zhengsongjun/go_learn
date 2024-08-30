package main

import (
	"errors"
	"fmt"
)

func read_file(file_name string) error {
	if file_name == "main.go" {
		return nil
	} else {
		return errors.New("我们需要一个main.go文件")
	}
}

func my_fun(file_name string) {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("给管理员发信息把")
			fmt.Printf("失败原因是 %v\n", err)
		}
	}()

	if err := read_file(file_name); err != nil {
		panic(err)
	} else {
		fmt.Println("文件读取成功")
	}
}

func main() {
	my_fun("main")
	fmt.Println("结束程序")
}
