package main

import (
	"fmt"
)

var done1 = make(chan bool)
var Msg1 string

func aGoroutine1() {

	Msg1 = "hello world"
	<-done1
	//close(done1)   // 等价于 done <- true
}

// channel 的简单示例
func main() {
	// 正常流程
	go aGoroutine1()

	done1 <- true
	fmt.Println(Msg1)
}
