package main

import "fmt"

type Person struct {
	name string
	age  int
	sex  string
}

// 给结构体加方法
func (p Person) PrintInfo() {
	fmt.Printf("自我介绍%v", p)
}

// 传入指针类型结构体
func (p *Person) setInfo(age int, sex string) {
	p.age = age
	p.sex = sex
}

// 结构体是值类型
func fn1() {
	p := Person{
		"zsj",
		20,
		"男",
	}
	p_copy := p
	fmt.Printf("p_copy: %v\n", p_copy)
	p_copy.age = 50
	p_copy.name = "test"
	fmt.Printf("p_copy: %v\n", p_copy)
	fmt.Printf("p: %v\n", p)
	p.PrintInfo()
	p_copy.PrintInfo()
}

func fn2() {
	p := Person{
		"test",
		100,
		"不男不女",
	}
	p.PrintInfo()
	p.setInfo(10, "男")
	p.PrintInfo()

}

func main() {
	fn1()
	fmt.Printf("\"____________________\": %v\n", "____________________")
	fn2()
}
