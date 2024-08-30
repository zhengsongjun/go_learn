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

func my_func(file_name string) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("读取文件失败，给管理员发个邮件")
			fmt.Printf("失败结果是 %v\n", err)
		}
	}()
	if err := read_file(file_name); err != nil {
		panic(err)
	} else {
		fmt.Println("读取成功")
	}
}

func main() {
	my_func("ss.go")
	fmt.Println("最后执行程序")
}
