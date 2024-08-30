package main

import "fmt"

func fn1() {
	fmt.Println("运行fn1")
}

// panic 可以单独使用，panic直接结束
// func fn_panic() {
// 	panic("第一个panic异常")
// }

// recover 必须 defer 要配合使用,但是不影响后面的函数执行
func fn2() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("err:", err)
		}
	}()
	panic("抛出一个异常")
}

func main() {
	// fn1()
	// fn_panic()
	// fmt.Printf("\"终止执行\": %v\n", "终止执行")

	fn1()
	fn2()
	fmt.Println("终止执行")
}
