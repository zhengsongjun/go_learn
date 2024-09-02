package main

import "fmt"

func read_file(file_name string) error {
	if file_name == "main.go" {
		return nil
	} else {
		return fmt.Errorf("你给错文件了，给的是%s，我们需要的是main.go", file_name)
	}
}

func fn1(file_name string) {

	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("err: %v\n", err)
		}
	}()

	if err := read_file(file_name); err != nil {
		panic(err)
	} else {
		fmt.Print("读取文件")
	}
}

func main() {
	fn1("test.go")
	fmt.Printf("\"______________\": %v\n", "______________")
	fn1("main.go")
}
