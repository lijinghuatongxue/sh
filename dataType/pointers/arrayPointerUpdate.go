package main

import "fmt"

func arrayPointerUpdate(arr *[3]int) {
	(*arr)[0] = 2
}

func slicePointerUpdate(slice []int) {
	slice[0] = 3
}

// 对函数内的数组进行一些修改
func main() {
	var a = [3]int{1, 2, 3}
	fmt.Println("array 数组 修改前 a: ", a)
	arrayPointerUpdate(&a)
	fmt.Println("array 数组 修改后 a: ", a)
	slicePointerUpdate(a[:])
	fmt.Println("slice 切片 修改后 a: ", a)

}
