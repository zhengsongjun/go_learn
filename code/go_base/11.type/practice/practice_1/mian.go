package main

import (
	"fmt"
)

type CustomInt int
type MyInt = int

func (c CustomInt) PrintInfo() {
	fmt.Printf("我的值是%v,我的类型是%T\n", c, c)
}

func customTypeAndTypeAlias() {
	var customInt CustomInt = 1
	var typeAlias MyInt = 1
	customInt.PrintInfo()
	fmt.Printf("typeAlias的值是: %v,它的类型是%T\n", typeAlias, typeAlias)
}

type PersonBase struct {
	name string
}

type Person struct {
	personBase PersonBase
	age        int
	address    Address
	hobby      []string
	bestFriend map[int]string
}

type Address struct {
	city        string
	street      string
	houseNumber int
}

func (p Person) PrintInfo() {
	fmt.Printf("你好，我叫%v,我的年龄是%v", p.personBase.name, p.age)
	for index, hobbyItem := range p.hobby {
		fmt.Printf("我的爱好%v%v", index+1, hobbyItem)
	}

	fmt.Printf("我家住在%v,%v,门牌号是%v", p.address.city, p.address.street, p.address.houseNumber)

}

func instantiateStruct() {
	p := Person{
		personBase: PersonBase{
			name: "test",
		},
		age: 20,
		address: Address{
			city:        "合肥",
			street:      "长宁大道",
			houseNumber: 30,
		},
		hobby: []string{
			"打篮球", "羽毛球", "乒乓球",
		},
		bestFriend: map[int]string{
			0: "testSelf",
			1: "wife",
			2: "parent",
		},
	}
	p.PrintInfo()
}

func main() {
	customTypeAndTypeAlias()
	fmt.Printf("\"_________________-\": %v\n", "_________________-")
	instantiateStruct()
}
