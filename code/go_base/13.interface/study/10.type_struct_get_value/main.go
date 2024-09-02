package main

import "fmt"

type Address struct {
	City        string
	Street      string
	HouseNumber string
}

func fn() {
	myMap := make(map[string]interface{})
	myMap["username"] = "test"
	myMap["address"] = Address{
		City:        "合肥",
		Street:      "长宁大道",
		HouseNumber: "100",
	}
	myMap["hobby"] = []string{"打篮球", "打羽毛球"}
	// myMap["hobby"][0] 报错，无法获取
	// myMap["address"].City 报错，无法获取
	if hobby, hobbyState := myMap["hobby"].([]string); hobbyState {
		fmt.Printf("hobby[1]%v\n", hobby[1])
	}

	if address, addressState := myMap["address"].(Address); addressState {
		fmt.Println(address.City)
	}

}

func main() {
	fn()
}
