package main

func main() {
	done := make(chan int)
	go func() {
		println("hello world")
		done <- 1
	}()
	<-done // 执行此行时，必须要求 done <-1 也已经执行，可以判定，当执行到此行时，print事件已经发生
}
