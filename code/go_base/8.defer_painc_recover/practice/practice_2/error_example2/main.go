package main

import (
	"fmt"
)

func read_file(file_name string) error {
	if file_name == "main.go" {
		return nil
	} else {
		return fmt.Errorf("我们需要的是main.go文件，而你的文件是%s\n", file_name)
	}
}

func fn1(file_name string) {

	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("err: %v\n", err)
			fmt.Printf("请联系管理员\n")
		}
	}()

	if err := read_file(file_name); err == nil {
		fmt.Print("文件读取成功\n")
	} else {
		fmt.Print(err)
		panic("文件读取错误")
	}

}

func main() {
	fn1("main.go")
	fmt.Printf("\"_________________\": %v\n", "_________________")
	fn1("test.go")
}
