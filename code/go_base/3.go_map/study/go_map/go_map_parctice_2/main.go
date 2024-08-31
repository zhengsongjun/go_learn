package main

import "fmt"

// 使用make创建map
func create_to_make_map() {
	my_map := make(map[string]string, 200)
	my_map["username"] = "test"
	my_map["age"] = "20"
	my_map["class"] = "10_2"
	my_map["sex"] = "男"
	for key, value := range my_map {
		fmt.Printf("key: %v ,value:%v --- ", key, value)
	}
}

func var_full_map() {
	my_map := map[string]string{
		"usename": "test",
		"age":     "20",
		"class":   "10_2",
		"sex":     "男",
	}
	fmt.Printf("my_map: %v\n", my_map)
}

func for_range_map() {
	my_map := map[string]string{
		"username": "test",
		"age":      "20",
		"class":    "10-2",
		"sex":      "男",
	}
	fmt.Printf("my_map: %v\n", my_map)

}

func find_item_in_map() {
	my_map := map[string]string{
		"username": "find_item_in_map",
		"age":      "20",
		"class":    "10-3",
	}
	_, state := my_map["username"]
	fmt.Printf("\"username 的值在my_map中存在吗？\": %v\n", state)
	_, state2 := my_map["username1"]
	fmt.Printf("\"username2的值在 my_map中存在吗？\": %v\n", state2)
}

func delete_map_item() {
	my_map := make(map[string]string)
	my_map["username"] = "test"
	my_map["age"] = "20"
	my_map["class"] = "10-2"
	my_map["sex"] = "男"
	fmt.Printf("my_map: %v\n", my_map)
	delete(my_map, "username")
	fmt.Printf("my_map: %v\n", my_map)
	delete(my_map, "age")
	fmt.Printf("my_map: %v\n", my_map)
	delete(my_map, "class")
	fmt.Printf("my_map: %v\n", my_map)
	delete(my_map, "sex")
	fmt.Printf("my_map: %v\n", my_map)
}

func main() {
	create_to_make_map()
	fmt.Printf("\"____________________-\": %v\n", "____________________-")
	var_full_map()
	fmt.Printf("\"---------------\": %v\n", "---------------")
	for_range_map()
	find_item_in_map()
	delete_map_item()
}
