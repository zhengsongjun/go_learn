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
	fmt.Printf("%v启动\n", p.Name)
}
func (p Phone) stop() {
	fmt.Printf("%v关闭\n", p.Name)
}

type Camcer struct {
	Name string
}

func (c Camcer) start() {
	fmt.Printf("%v启动\n", c.Name)
}

func (c Camcer) stop() {
	fmt.Printf("%v关闭\n", c.Name)
}

func startAndStop(u Usber) {
	u.start()
	u.stop()
}

func main() {
	phone := Phone{
		Name: "小米",
	}
	camcer := Camcer{
		Name: "小米",
	}

	startAndStop(phone)
	startAndStop(camcer)

}
