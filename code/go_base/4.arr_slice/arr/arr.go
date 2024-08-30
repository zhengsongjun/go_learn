package main;

import "fmt";


func main() {
	// 第一种
	// var arr [5]int
	// arr[0] = 0;
	// arr[1] = 1;
	// arr[2] = 2;
	// arr[3] = 3;
	// arr[4] = 4;
	// 直接 第二种
	// var arr = [5]int{1,2,3,4,5}
	// 第三种
	arr := [...]int{1,2,3,4,5,6}
	for i  := 0 ; i < len(arr) ; i++ {
		fmt.Println(arr[i])
	}
}