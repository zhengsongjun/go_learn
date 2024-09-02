package main

import "fmt"

// 匿名字段不允许重复
type Person struct {
	string
	int
}

func main() {
	person := Person{
		"test",
		20,
	}

	fmt.Printf("person: %v\n", person)
	fmt.Printf("person: %#v\n", person)
}
