package main

import "fmt"

func arrayUpdate(arr *[3]int) {
	(*arr)[0] = 2
}

// 对函数内的数组进行一些修改
func main() {
	var a = [3]int{1, 2, 3}
	fmt.Println("修改前 a: ", a)
	arrayUpdate(&a)
	fmt.Println("修改后 a: ", a)
}
