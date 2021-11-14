package main

import "fmt"

func main() {
	var done = make(chan int, 3)
	go func() {
		fmt.Println("hello world!")
		done <- 1 // 在有缓存的channel中，why？ todo
	}()
	<-done
}
