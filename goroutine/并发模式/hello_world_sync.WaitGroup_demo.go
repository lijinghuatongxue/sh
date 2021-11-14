package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	// go func 开启多个线程(goroutine)去执行
	for i := 0; i < 10; i++ {
		// wg.Add(1) 用于增加等待事件的个数，必须保证在后台线程开启之前执行
		wg.Add(1)
		go func() {
			fmt.Println("hello world!")
			// wg.Done() 表示完成一个事件
			wg.Done()
		}()
	}
	// 等待所有的线程(事件)完成
	wg.Wait()
}
