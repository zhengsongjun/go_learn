package main

import "fmt"

type Person struct {
	Name    string
	Age     int
	Address Address
}

type Address struct {
	City        string
	Street      string
	HouseNumber int
}

func main() {
	p := Person{
		Name: "test",
		Age:  20,
		Address: Address{
			City:        "合肥",
			Street:      "长宁大道",
			HouseNumber: 20,
		},
	}
	fmt.Printf("p: %#v\n", p)
}
