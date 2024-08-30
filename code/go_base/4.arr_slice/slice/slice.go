package main

import "fmt"

func slice_with_arr() {
	// arr := []string { "java", "php" , "nodejs" , "go"}
	// for k,v := range arr {
	// 	fmt.Printf("%v : %v",k ,v);
	// }
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	// for k,v := range arr {
	// 	fmt.Printf("%v:%v\n",k,v)
	// }

	slice := arr[3:]
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice)) // 切片基于数组的话，那么切片的切的数组的第一个元素到数组最后一个元素的长度就是切片的容量cap
}

func extend_slice() {
	slice := []int{}
	// slice[0] = 1 这么做是错的，因为不允许在未知的长度做操作
	slice = append(slice, 1)
	fmt.Println(slice)
}

func merge_slice() {
	sliceA := []int{1, 2, 3, 4}
	sliceB := []int{5, 6, 7, 8, 9}
	sliceA = append(sliceA, sliceB...)
	fmt.Println(sliceA) // [1 2 3 4 5 6 7 8 9]
}

func capacity_slice_strategy() {
	sliceA := []int{1}
	for i := 0; len(sliceA) < 100000; i++ {
		sliceA = append(sliceA, sliceA...)
		fmt.Printf("slice 长度%v,容量%v\n", len(sliceA), cap(sliceA))
	}
}

func copy_slice() {
	sliceA := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	sliceB := make([]int, len(sliceA))
	copy(sliceB, sliceA)
	sliceB = append(sliceB, 10)
	fmt.Println(sliceA)
	fmt.Println(sliceB)
}

func delete_slice_item(i int, slice []int) []int {
	if i < 0 || i >= len(slice) {
		return slice
	}

	slice = append(slice[0:i], slice[i+1:]...)
	return slice
}

func main() {
	// slice_with_arr();

	// 切片扩容
	// extend_slice();
	// 合并分片
	// merge_slice();
	// 切片的扩容策略
	// capacity_slice_strategy();
	// 切片的复制
	// copy_slice();

	// 通过append删除切片元素
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(delete_slice_item(2, slice))
}
