package main

import "fmt"
func slice_with_arr () {
	arr := [10]int{1,2,3,4,5,6,7,8,9,10}
	slice := arr[3:4]
	fmt.Println(arr)
	fmt.Println(slice)
	fmt.Println(len(slice)) // 1 
	fmt.Println(cap(slice)) // 7

}

func extend_slice(){
	slice := []int{1,2,3,4,5}
	fmt.Println(slice)
	fmt.Println(append(slice,6))	
}

func merge_slice(){
	sliceA := []int{1,2,3,4,5}
	sliceB := []int{6,7,8,9,10}
	sliceB = append(sliceA,sliceB...)
	fmt.Println(sliceB)
}

func copy_slice(){
	sliceA := []int{1,2,3,4,5}
	sliceB := make([]int,len(sliceA))
	copy(sliceB,sliceA)
	sliceB[0] = 2
	fmt.Println(sliceA)
	fmt.Println(sliceB)
}

func delete_slice_to_index(i int,slice []int)[]int{
	if i < 0 || i >= len(slice) {
		return slice
	}

	return append(slice[0:i],slice[i+1:]...)
}

func main () {
	//slice_with_arr();
	// extend_slice()
	// merge_slice()
	// copy_slice()
	slice := []int{1,2,3,4,5,6,7,8,9,10}
	fmt.Println(delete_slice_to_index(5,slice))
}