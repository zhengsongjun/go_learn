package main

import (
	"errors"
	"fmt"
)

func read_file(file_name string) error {
	if file_name == "main.go" {
		return nil
	} else {
		return errors.New("我们需要的事main")
	}
}

func my_fn(file_name string) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			fmt.Println("请联系管理员更换文件")
		}
	}()
	if err := read_file(file_name); err != nil {
		panic(err)
	} else {
		fmt.Println("读取成功")
	}
}

func main() {
	my_fn("main.go")
	fmt.Println("最后执行")
}
