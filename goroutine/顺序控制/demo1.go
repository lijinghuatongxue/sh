package main

import (
	"fmt"
	"sync"
)

func main() {
	var mu sync.Mutex
	go func() {
		fmt.Println("hello world")
		mu.Unlock()
	}()
	mu.Lock() // 执行到此行时，mu.Unlock() 已经执行，因为同一个线程的顺序可以保证，所有此时打印事件已经发生
}
