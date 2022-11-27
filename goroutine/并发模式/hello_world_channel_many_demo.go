package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"time"
)

func main() {
	var done = make(chan int, 10)
	logrus.Infof("chal.cap.is:%v", cap(done))
	for i := 0; i < cap(done); i++ {
		go func() {
			// fmt
			fmt.Println("hello world")
		}()
		done <- 1
	}
	for i := 0; i < cap(done); i++ {
		time.Sleep(1 * time.Second) // 这里出现等待，两个for同时运行，没有将通道消费完，程序不停止
		<-done
	}
}
