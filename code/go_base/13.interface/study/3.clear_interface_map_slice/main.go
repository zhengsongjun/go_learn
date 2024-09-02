package main

import "fmt"

func clear_interface_map() {
	myMap := make(map[string]interface{})
	myMap["username"] = "test"
	myMap["age"] = 12
	myMap["isLiked"] = true
	fmt.Printf("myMap: %v\n", myMap)
}

func clear_interface_slice() {
	mySlice := make([]interface{}, 0)
	mySlice = append(mySlice, "test")
	mySlice = append(mySlice, 20)
	mySlice = append(mySlice, true)
	fmt.Printf("mySlice: %v\n", mySlice)
}

func main() {
	clear_interface_map()
	clear_interface_slice()
}
