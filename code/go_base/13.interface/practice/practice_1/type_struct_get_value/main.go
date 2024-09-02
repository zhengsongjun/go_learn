package main

import "fmt"

type Address struct {
	City        string
	Street      string
	HouseNumber int
}

func fn1() {
	myMap := make(map[string]interface{})
	myMap["name"] = "test"
	myMap["address"] = Address{
		City:        "北京",
		Street:      "长安路",
		HouseNumber: 100,
	}
	myMap["hobby"] = []string{"打篮球", "打乒乓球", "踢足球"}
	if address, ok := myMap["address"].(Address); ok {
		fmt.Printf("address的城市名: %v\n", address.City)
	}
	if hobby, ok := myMap["hobby"].([]string); ok {
		fmt.Printf("hobby[1]: %v\n", hobby[1])
	}
}

func main() {
	fn1()
}
