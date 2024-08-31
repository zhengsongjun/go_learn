package main

import "fmt"

// func create_to_make_map() {
// 	myMap := make(map[string]string, 10)
// 	myMap["username"] = "test"
// 	myMap["age"] = "20"
// 	myMap["class"] = "10-2"

// 	for key, value := range myMap {
// 		fmt.Println(key, value, "\n")
// 	}

// }

// func var_full_map() {
// 	myMap := map[string]string{
// 		"username": "zhengsongjun",
// 		"age":      "20",
// 		"class":    "10_2",
// 		"sex":      "男",
// 	}

// 	for key, value := range myMap {
// 		fmt.Println(key, value, "\n")
// 	}
// }

// func for_range_map() {
// 	myMap := make(map[string]string, 10)
// 	myMap["username"] = "test"
// 	myMap["age"] = "20"
// 	myMap["class"] = "10-2"
// 	for key, value := range myMap {
// 		fmt.Println(key, value, "\n")
// 	}
// }

// func find_item_in_map() {
// 	myMap := map[string]string{
// 		"username": "test",
// 		"age":      "20",
// 		"class":    "10-2",
// 	}

// 	if _, state := myMap["username"]; state {
// 		fmt.Printf("username在myMap有值吗？%v \n", state)
// 	}

// 	if _, state1 := myMap["test"]; state1 {
// 		fmt.Printf("test在myMap有值吗 %v \n", state1)
// 	}

// }

func delete_map_item() {
	myMap := map[string]string{
		"username": "test",
		"age":      "20",
	}

	delete(myMap, "username")
	fmt.Printf("myMap: %v\n", myMap)

}

func main() {
	// create_to_make_map()
	// var_full_map()
	// for_range_map()
	// find_item_in_map()
	delete_map_item()
}
