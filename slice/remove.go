package main

import "fmt"

func main()  {
	a := []int{1,2,3,4,5,6,7,8,9}
	// ==================================== remove ====================================

	// 尾部删除,删除最后一个元素
	a = a[:len(a)-1]
	fmt.Println(a)
	// 尾部删除,删除最后N个元素
	n := 1
	a = a[:len(a)-n]
	fmt.Println(a)
	// 头部元素删除，删除第一个元素
	a = a[1:]
	fmt.Println(a)
	// 头部元素删除，删除开头N个元素
	a = a[n:]
	fmt.Println(a)

	// 不会产生内存空间移动的删除 append
	// 头部元素删除，删除第一个元素，删除就是增加
	a = append(a[:0],a[1:]...)
	fmt.Println(a)
	// 头部元素删除，删除第N个元素，删除就是增加
	a = append(a[:0],a[n:]...)
	fmt.Println(a)

	// copy 删除
	// 头部元素删除，删除开头第一个元素，先运行copy，使用copy得出复制的元素个数，该个数+1就是原来元素的len，此时a已经删除了第一个元素
	// copy return : returns the number of elements copied, which will be the minimum of
	a = a[:copy(a,a[1:])]
	fmt.Println(a)
	// 头部元素删除，删除开头N个元素
	a = a[:copy(a,a[n:])]
	fmt.Println(a)
}
