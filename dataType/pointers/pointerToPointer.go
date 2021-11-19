package main

import "fmt"

// 指针的指针
func main() {
	// 定义a 为 int
	var a int
	a = 67
	// 定义第一个指针
	var ptrOne *int
	// 定义第二个指针 **type
	var ptrSecond **int
	ptrOne = &a
	ptrSecond = &ptrOne
	// 打印指针的指针的值
	fmt.Println(**ptrSecond)
}
