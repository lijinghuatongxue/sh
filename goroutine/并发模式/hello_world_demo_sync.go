package main

import (
	"fmt"
	"sync"
	"time"
)

// 原理：先加锁，在解锁，当前加锁状态下，mu.lock 会阻塞，类似于 select{}
func main() {
	var mu sync.Mutex
	mu.Lock() // 1. 先锁
	go func() {
		fmt.Println("hello world!")
		time.Sleep(3 * time.Second)
		mu.Unlock()
	}()
	mu.Lock() // 2. 未解锁的话，会产生阻塞，直到解锁，如果程序执行到这里，也就意味这Println事件已经发生
}
