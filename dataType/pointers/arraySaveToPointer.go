package main

import "fmt"

// 用指针保存一个数组
func main() {
	var a = [3]int{1, 23, 67}
	const MAX int = 3
	var i int
	var ptr [MAX]*int
	for i = 0; i < 3; i++ {
		ptr[i] = &a[i]
	}
	for i = 0; i < 3; i++ {
		fmt.Printf("指针 %d | 值为 %d \n", i, *ptr[i])
	}
}
