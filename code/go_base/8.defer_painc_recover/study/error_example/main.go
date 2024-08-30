package main

import (
	"errors"
	"fmt"
)

func readFile(fileName string) error {
	if fileName == "main.go" {
		return nil
	} else {
		return errors.New("读取文件失败")
	}
}

func my_fn() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("err: %v\n", err)
			fmt.Printf("\"给管理员发送邮件\": %v\n", "给管理员发送邮件")
		}
	}()
	if err := readFile("test.go"); err != nil {
		panic(err)
	}
}

func main() {
	my_fn()
	fmt.Println("最后执行")
}
