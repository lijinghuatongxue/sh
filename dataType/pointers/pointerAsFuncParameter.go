package main

import "fmt"

func main() {
	var a int = 3000
	var b int = 3000
	fmt.Printf("修改前 | a: %d  |b:%d \n", a, b)
	// 传入指针参数
	swap(&a, &b)
	fmt.Printf("修改后 | a: %d  |b:%d \n", a, b)

}

func swap(pointerA *int, pointerB *int) {
	var temp int = 67
	*pointerA = temp      // 保存 temp（67 int） 的值到指针A
	*pointerB = *pointerA // 赋值 A (67 int) 指针的值到 指针 B
	*pointerB = temp      // 赋值 temp （67 int） 的值到指针B
}
