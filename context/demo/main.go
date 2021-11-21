package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// 初始化一个context，超时时间设置为1s
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	// 并发一个 goroutine
	go handle(ctx, 500*time.Millisecond)
	select {
	// ctx.done，本质是一个channel，只在整个context执行完成或者被取消后才会返回，多次调用done方法，只会返回同一个channel
	// 本次handle执行只会耗时500ms，所以本次case不会被执行
	case <-ctx.Done():
		fmt.Println("main", ctx.Err())
	}
}

func handle(ctx context.Context, duration time.Duration) {
	var c = make(chan int, 1)
	select {
	// 该条件不满足
	case <-ctx.Done():
		fmt.Println("handle", ctx.Err())
	// 延迟500ms，该条件可以满足
	case <-time.After(duration):
		fmt.Println("process request with", duration)
	// 以下两个case随机选择，在channel 缓存为1的情况下
	case c <- 1:
		time.After(230 * time.Millisecond)
		fmt.Println("耗时230ms")
	case c <- 1:
		time.After(190 * time.Millisecond)
		fmt.Println("耗时190ms")
	default:
		fmt.Println("我是永远不会被执行的")
	}

}
