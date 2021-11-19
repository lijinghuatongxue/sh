package main

import "fmt"

func main() {
	var a int = 20 /* 声明实际变量 */
	var ip *int    /* 声明指针变量 */

	ip = &a /* 指针变量的存储地址 */

	fmt.Printf("a 变量的地址是: %x\n", &a)

	/* 指针变量的存储地址 */
	fmt.Printf("ip 变量的存储地址: %x\n", ip)

	/* 使用指针访问值，获取指针的值 */
	fmt.Printf("*ip 变量的值: %d\n", *ip)

	// 改变指针的值, 原变量也会跟着改变
	*ip = *ip + 1
	fmt.Printf("*ip 改变后的指针的值: %d\n", *ip)
	fmt.Printf("a   改变后原变量的值: %d\n", a)
}
