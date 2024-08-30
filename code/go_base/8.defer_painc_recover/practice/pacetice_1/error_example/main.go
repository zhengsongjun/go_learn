package main

import (
	"errors"
	"fmt"
)

func read_file(filename string) error {
	if filename != "main.go" {
		return errors.New("函数需要main.go文件")
	} else {
		return nil
	}
}

func my_fn() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			fmt.Println("后面不用处理了，给管理员发个邮件")
		}
	}()
	if err := read_file("test.go"); err != nil {
		panic("读取失败")
	} else {
		fmt.Println("读取成功")
	}
}

func main() {
	my_fn()
	fmt.Println("程序结束")
}
