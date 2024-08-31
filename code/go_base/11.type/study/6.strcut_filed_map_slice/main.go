package main

import "fmt"

type Person struct {
	name       string
	age        int
	hobby      []string
	bestFriend map[int]string
}

func personInfo() {
	zsj := Person{
		name: "zsj",
		age:  20,
	}
	zsj.hobby = []string{"打篮球", "打羽毛球"}
	zsj.bestFriend = map[int]string{
		0: "myselft",
		1: "wife",
		2: "parent",
	}

	fmt.Printf("zsj: %#v\n", zsj)
}

func main() {
	personInfo()
}
