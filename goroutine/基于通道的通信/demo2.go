package main

import (
	"github.com/sirupsen/logrus"
)

// 控制通道的缓存大小 page48
func main() {
	var limit = make(chan int, 10)
	c := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for _, w := range c {
		w := w
		go func() {
			limit <- 1
			logrus.Info(w)
			<-limit
		}()
	}
	//for _,w := range c{
	//	logrus.Warn(w)
	//}
	select {}
}
