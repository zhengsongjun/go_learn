package main

import "fmt"

// 使用make创建map
func create_to_make_map() {
	my_map := make(map[int]string)
	my_map[0] = "0"
	my_map[1] = "1"
	my_map[2] = "2"
	for k, v := range my_map {
		fmt.Printf("k: %v, v: %v \n", k, v)
	}
}

func var_full_map() {
	my_map := map[string]string{
		"username": "zhangsan",
		"age":      "20",
		"class":    "10-2",
	}

	for k, v := range my_map {
		fmt.Printf("k: %v , i:%v\n", k, v)
	}
}

func for_range_map() {
	my_map := map[string]string{
		"usename": "kangjie",
		"age":     "20",
		"class":   "class",
	}
	for key, value := range my_map {
		fmt.Printf("key: %v , value : %v\n", key, value)
	}
}

func find_item_in_map() {
	my_map := map[string]string{
		"username": "test",
		"age":      "20",
		"class":    "10-2",
	}

	_, state := my_map["username"]
	fmt.Printf("\"username的值在my_map里面存在吗？ \": %v\n", state)
	_, state2 := my_map["age2"]
	fmt.Printf("\"age2在my_map里面存在吗\": %v\n", state2)

}

func delete_map_item() {
	my_map := make(map[string]string)
	my_map["usename"] = "test"
	my_map["age"] = "20"
	my_map["class"] = "10_2"
	fmt.Printf("my_map: %v\n", my_map)
	delete(my_map, "usename")
	delete(my_map, "age")
	delete(my_map, "class")
	fmt.Printf("my_map: %v\n", my_map)
}

func main() {
	create_to_make_map()
	var_full_map()
	fmt.Printf("\"---------------\": %v\n", "---------------")
	for_range_map()
	find_item_in_map()
	delete_map_item()
}
