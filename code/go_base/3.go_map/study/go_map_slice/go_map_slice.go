package main
import (
	"fmt"
	"sort"
)

func slice_map () {
	slice := make([]map[string]string,0);
	slice = append(slice,map[string]string{
		"usename":"test1",
		"age" : "21",
		"sex" : "男",
	})

	slice = append(slice,map[string]string{
		"usename":"test2",
		"age" : "22",
		"sex" : "男1",
	})

	slice = append(slice,map[string]string{
		"usename":"test3",
		"age" : "23",
		"sex" : "男2",
	})
	for k,v := range slice {
		fmt.Println(k,v)
	}
}

func map_slice () {
	map_slice := make(map[string][]string,0)
	map_slice["hobby"] = []string{
		"吃饭","睡觉",
	}
	map_slice["language"] = []string{
		"php","java",
	}
	fmt.Println(map_slice)
}

func map_slice_sort_by_key () {
	map_slice := make(map[int]int)
	map_slice[0] = 10
	map_slice[3] = 900
	map_slice[5] = 88
	map_slice[10] = 0
	map_slice[100] = 11

	slice := make([]int,0)
	for k := range map_slice {
		slice = append(slice,k)	
	}
	sort.Ints(slice)
	for _,v := range slice {
		fmt.Println(map_slice[v])
	}
}

func map_slice_sort_by_value () {
	map_slice := make(map[int]int)
	map_slice[0] = 0
	map_slice[1] = 10
	map_slice[2] = 9
	map_slice[3] = 3
	map_slice[4] = 12

	fmt.Println(len(map_slice))

	slice := make([]int,len(map_slice))
	i := 0
	for _,v := range map_slice {
		slice[i] = v
		i++
	}
	sort.Ints(slice)
	fmt.Println(slice)
}

func main () {
	// slice_map()
	// map_slice()
	// map_slice_sort_by_key()
	map_slice_sort_by_value()
}