package main

import "fmt"

type Usbr interface {
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

func (c Camcer) start() {
	fmt.Println("相机开机")
}

func (c Camcer) stop() {
	fmt.Println("相机关机")
}

func stopOrStart(u Usbr) {
	if _, ok := u.(Phone); ok {
		fmt.Println("我检测到了，你是手机，所以我不能给你开机，你关机把\n")
		u.stop()
	}

	if _, ok := u.(Camcer); ok {
		fmt.Println("我检到了，你是相机，你开机运行把\n")
		u.start()
	}
}

func main() {

	phone := Phone{
		Name: "小米手机",
	}

	camcer := Camcer{
		Name: "华为相机",
	}

	stopOrStart(phone)
	stopOrStart(camcer)

}
