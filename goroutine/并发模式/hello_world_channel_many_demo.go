package main

import "fmt"

func main() {
	var done = make(chan int, 10)
	for i := 0; i < cap(done); i++ {
		go func() {
			// fmt
			fmt.Println("hello world")
		}()
		done <- 1
	}
	for i := 0; i < cap(done); i++ {
		<-done
	}
}
