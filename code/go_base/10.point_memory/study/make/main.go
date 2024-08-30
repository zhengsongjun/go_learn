package main

import "fmt"

// 切片和map必须通过make创建了内存空间之后，才能进行操作添加

// func no_make() {
// 	var slice []int
// 	slice[0] = 1 // 报错  index of nil slice

// 	var my_map map[string]string
// 	my_map["name"] = "test" // nil derefence in map update
//  var i *int
//  i = 10  // 报错，指针也是引用数据类型
// }

func has_make() {
	slice := make([]int, 10)
	slice[0] = 10
	fmt.Printf("slice: %v\n", slice)
	my_map := make(map[string]string, 10)
	my_map["name"] = "test"
	fmt.Print(my_map)
}

func main() {
	has_make()
}
