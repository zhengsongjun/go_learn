package main

import "fmt"

// 变量，结构体，函数首字母小写，只能这个包内用
// 变量，结构体，函数首字母大写，其他包也可以使用
// 包括结构体的字段名首字母
type Person struct {
	name string
	age  int
	sex  string
}

// 结构体实例化 第一种
func instantiatePersonStruct1() {
	var p Person
	p.name = "张三"
	p.age = 20
	p.sex = "男"
	fmt.Printf("p: %v\n", p)
	fmt.Printf("p: %T\n", p)
	fmt.Printf("p: %#v\n", p)
}

// 结构体实例化 第二种

func instantiatePersonStruct2() {
	var p = new(Person)
	p.name = "李四"
	p.age = 18
	p.sex = "女"
	fmt.Printf("Person: %v\n", p)
	fmt.Printf("person: %T\n", p)
	fmt.Printf("person: %#v\n", p)
	// golang 中支持对结构体指针直接使用.来访问结构体的成员，p.name = "李四" 其实底层就是(*p).name = "李四"
	(*p).name = "李四，指针修改"
	fmt.Printf("p: %v\n", p)
}

func instantiatePersonStruct3() {
	p3 := &Person{}
	p3.name = "王五"
	p3.age = 19
	p3.sex = "不男不女"
	fmt.Printf("p3: %v\n", p3)
}

func instantiatePersonStruct4() {
	p := Person{
		name: "赵六",
		age:  100,
		sex:  "非男肥女",
	}
	fmt.Printf("p: %v\n", p)

}

func instantiatePersonStruct5() {
	p := &Person{
		name: "王麻子",
		age:  3,
		sex:  "人妖",
	}
	fmt.Printf("p: %v\n", p)
}

func instantiatePersonStruct6() {
	p := &Person{
		"王麻子",
		3,
		"人妖",
	}
	fmt.Printf("p: %v\n", p)
}

func main() {
	instantiatePersonStruct1()
	fmt.Printf("\"____________________\": %v\n", "____________________")
	instantiatePersonStruct2()
	fmt.Printf("\"____________________\": %v\n", "____________________")
	instantiatePersonStruct3()
	fmt.Printf("\"____________________\": %v\n", "____________________")
	instantiatePersonStruct4()
	fmt.Printf("\"____________________\": %v\n", "____________________")
	instantiatePersonStruct5()
	fmt.Printf("\"____________________\": %v\n", "____________________")
	instantiatePersonStruct6()
}
