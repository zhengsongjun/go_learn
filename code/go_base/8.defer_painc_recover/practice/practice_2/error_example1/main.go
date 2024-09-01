package main

import (
	"errors"
	"fmt"
)

func read_file(file_name string) error {
	if file_name == "main.go" {
		fmt.Print("读取main.go文件\n")
		return nil
	} else {
		return errors.New("文件错误不是我们需要的文件")
	}
}

func fn1(file_name string) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Printf("err: %v\n", err)
			fmt.Print("联系管理员发送正确文件\n")
		}
	}()

	err := read_file(file_name)
	if err != nil {
		panic(err)
	} else {
		fmt.Print("文件读取成功")
	}

}

func main() {
	fn1("test.go")
	fmt.Printf("\"________________\": %v\n", "________________")
	fn1("main.go")
}
