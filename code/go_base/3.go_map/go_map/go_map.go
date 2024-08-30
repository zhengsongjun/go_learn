package main
import "fmt"

// 使用make创建map
func create_to_make_map () {
	useinfo := make(map[string]string,200)
	useinfo["username"] = "张三"
	useinfo["age"] = "20"
	useinfo["sex"] = "男"
	fmt.Println(useinfo)
	fmt.Println(useinfo["username"])
	fmt.Println(useinfo["age"])
	fmt.Println(useinfo["sex"])
}

func var_full_map () {
	useinfo := map[string]string{
		"username" : "张三",
		"age"  	   : "20",
		"sex"      : "男",
	}
	fmt.Println(useinfo)
	fmt.Println(useinfo["username"])
	fmt.Println(useinfo["age"])
	fmt.Println(useinfo["sex"])
}

func for_range_map () {
	user_info := map[string]string{
		"username" : "张三",
		"age"      : "20",
		"sex"      : "男",
	}
	for k,v := range user_info {
		fmt.Printf("key:%v,value:%v\n",k,v)
	}
}

func find_item_in_map() {
	user_info := map[string]string{
		"user_name" : "张三",
		"age" : "20",
		"sex" : "男",
	}
	v,state := user_info["user_name"]
	fmt.Println(v,state) // 张三 true
	v1,state1 := user_info["xxxxx"]
	fmt.Println(v1,state1) // 空 false
}

func delete_map_item () {
	user_info := map[string]string{
		"user_name" : "张三",
		"age" : "20",
		"sex" : "男",
	}
	fmt.Println(user_info)
	delete(user_info,"sex")
	delete(user_info,"user_name")
	fmt.Println(user_info)
}


func main() {
	// create_to_make_map();
	// var_full_map()
	// for_range_map()
	// find_item_in_map()
	delete_map_item()

}