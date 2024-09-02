package main

import "fmt"

func clear_interface_assert() {
	var a interface{}
	a = "string"
	if type0, ok := a.(string); ok {
		fmt.Printf("a的类型是%T\n", type0)
	}

	var b interface{}
	b = 100
	if type1, ok1 := b.(int); ok1 {
		fmt.Printf("b的类型是%T\n", type1)
	}
}

func toClearInterfaceAssert(a interface{}) {
	if _, ok := a.(string); ok {
		fmt.Println("a的类型是string")
	} else if _, ok := a.(int); ok {
		fmt.Println("a的类型是int")
	} else if _, ok := a.(bool); ok {
		fmt.Println("a的类型是bool")
	} else {
		fmt.Println("我也不知道a的类型")
	}
}

func toClearInterfaceAssert2(a interface{}) {
	switch a.(type) {
	case string:
		fmt.Println("a的类型是string!")
	case int:
		fmt.Println("a的类型是int!")
	case bool:
		fmt.Println("a的类型是布尔!")
	default:
		fmt.Println("我也不知道a的类型!")
	}
}

func main() {
	clear_interface_assert()

	toClearInterfaceAssert("100")
	toClearInterfaceAssert(false)
	toClearInterfaceAssert(100)
	fmt.Printf("\"______________________\": %v\n", "______________________")
	toClearInterfaceAssert2("100")
	toClearInterfaceAssert2(false)
	toClearInterfaceAssert2(100)
}
