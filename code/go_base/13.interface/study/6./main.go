// 结构体中值接受者，指针接受者实现接口的区别
package main

import "fmt"

type Usber interface {
	start()
	stop()
}

type Phone struct {
	Name string
}

func (p Phone) start() {
	fmt.Println("手机开机")
}
func (p Phone) stop() {
	fmt.Println("手机关机")
}

type Camcer struct {
	Name string
}

func (c *Camcer) start() {
	fmt.Println("相机开机")
}

func (c *Camcer) stop() {
	fmt.Println("相机关机")
}

func main() {
	var phone Phone = Phone{
		Name: "小米手机",
	}
	var p Usber
	p = phone
	p.start()

	camcer := Camcer{
		Name: "华为相机",
	}
	p = &camcer
	p.start()
}
