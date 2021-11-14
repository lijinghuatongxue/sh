package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Producer(factor int, out chan<- int) {
	for i := 0; ; i++ {
		time.Sleep(2 * time.Second)
		out <- i * factor
		fmt.Printf("生产了 %d\n", i*factor)
	}
}

func Consumer(in <-chan int) {
	for v := range in {
		time.Sleep(2 * time.Second)
		fmt.Printf("消费到 %d，当前队列长度 %d\n", v, len(in))
	}
}

func main() {
	ch := make(chan int, 64)
	go Producer(3, ch)
	go Producer(5, ch)
	go Consumer(ch)

	// ctrl + c 才能退出
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTRAP)
	fmt.Printf("quit (%v)\n", <-sig)
	//time.Sleep(3*time.Second)
}
