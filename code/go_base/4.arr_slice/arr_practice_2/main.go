package main

// func slice_with_arr() {
// 	arr := [5]int{1, 2, 3, 4, 5}
// 	sliceA := arr[:3]
// 	sliceB := arr[3:5]
// 	fmt.Println(sliceA)
// 	fmt.Println(sliceB)
// }

// func extend_slice() {
// 	slice := []int{1, 2, 3, 4, 5, 6}
// 	slice = append(slice, 7)
// 	slice = append(slice, 8)
// 	slice = append(slice, 9)
// 	slice = append(slice, 10)
// 	fmt.Println(slice)
// }

// func merge_slice() {
// 	sliceA := []int{1, 2, 3, 4, 5}
// 	sliceB := []int{6, 7, 8, 9, 10}
// 	sliceB = append(sliceA, sliceB...)
// 	for _, i := range sliceB {
// 		fmt.Println(i)
// 	}
// }

// func copy_slice() {
// 	sliceA := []int{1, 2, 3, 4, 5}
// 	sliceB := make([]int, len(sliceA))
// 	copy(sliceB, sliceA)
// 	fmt.Println(sliceA, sliceB)
// 	for i, v := range sliceB {
// 		sliceB[i] = v * 2
// 	}
// 	fmt.Printf("sliceA: %v\n", sliceA)
// 	fmt.Println(sliceB)
// }

// func delete_slice_to_index(i int, slice []int) []int {
// 	if i < 0 || i > len(slice) {
// 		return slice
// 	}
// 	return append(slice[:i], slice[i+1:]...)
// }

func main() {
	// slice_with_arr()
	// extend_slice()
	// copy_slice()
	//slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// fmt.Printf("delete_slice_to_index(2, slice): %v\n", delete_slice_to_index(2, slice))
}
