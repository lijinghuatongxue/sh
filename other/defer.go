package main

import (
	"fmt"
	"time"
)

func main() {
	startedAt := time.Now()
	defer func() { fmt.Println(time.Since(startedAt)) }()
	defer test() // 后进先出
	time.Sleep(time.Second * 2)
}
func test() {
	startedAt := time.Now()
	defer func() { fmt.Println(time.Since(startedAt)) }()
	time.Sleep(1 * time.Second)
}
