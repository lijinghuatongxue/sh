package main

import (
	"fmt"
)

var done = make(chan bool)
var Msg string

func aGoroutine() {

	Msg = "hello world"
	//done <- true
	close(done) // 等价于 done <- true
}

// channel 的简单示例
func main() {
	// 正常流程
	go aGoroutine()
	<-done
	fmt.Println(Msg)
}
