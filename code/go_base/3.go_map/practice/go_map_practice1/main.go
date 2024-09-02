package main

import "fmt"

// 使用make创建map
// func create_to_make_map () {
// 	my_map := make(map[string]string,10)
// 	my_map["name"] = "test"
// 	my_map["age"] = "20"
// 	fmt.Println(my_map)
// }

// func var_full_map () {
// 	my_map := map[string]string{
// 		"username":"test",
// 		"age":"20",
// 	}
// 	for key, value := range my_map {
// 		fmt.Println(key,value,"\n")
// 	}

// }

// func for_range_map () {
// 	my_map := map[string]string{
// 		"username":"test",
// 		"age":"20",
// 	}
// 	for key,value := range my_map {
// 		fmt.Println(key,value,"\n")
// 	}
// }

// func find_item_in_map() {
// 	myMap := map[string]string{
// 		"username": "test",
// 		"age":      "20",
// 	}
// 	if username, state := myMap["username"]; state {
// 		fmt.Printf("username有,是 %v\n", username)
// 	} else {
// 		fmt.Printf("username没有\n")
// 	}

// 	if test, state1 := myMap["test"]; state1 {
// 		fmt.Printf("test的值有，是%v\n", test)
// 	} else {
// 		fmt.Printf("test的值没有\n")
// 	}
// }

func delete_map_item() {
	myMap := map[string]string{
		"username": "test",
		"age":      "20",
	}

	delete(myMap, "username")
	for key, value := range myMap {
		fmt.Println(key, value)
	}
}

func main() {
	// create_to_make_map();
	// var_full_map()
	// for_range_map()
	// find_item_in_map()
	delete_map_item()

}
