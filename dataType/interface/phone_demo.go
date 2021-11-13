package main

import (
	"fmt"
)

// USB 声明一个USB的接口
type USB interface {
	Name() string
	Connect() // embed Connect方法，去获得fmt能力
}

// PhoneConnecter 声明一个PhoneConnect的结构去实现USB接口
type PhoneConnecter struct {
	name string
}

func (pc PhoneConnecter) Name() string {
	fmt.Sprintf("pc name is %s", pc.name)
	return pc.name
}

func (pc PhoneConnecter) Connect() {
	fmt.Println(pc.Name(), "is connected..")
}

func main() {
	var a USB                            //声明a的类型是USB类型
	a = PhoneConnecter{"you's computer"} //成功实现USB接口
	fmt.Println(a.Name())
	a.Connect()

}
