package main

import (
	"fmt"
	"reflect"
)

func main() {
	// 数组是值类型 Go中的数组是值类型，而不是引用类型。这意味着当它们被分配给一个新变量时，
	// 将把原始数组的副本分配给新变量。如果对新变量进行了更改，则不会在原始数组中反映。
	a := [...]string{"67", "jinghua", "meihua", "233", "267"}
	b := a      // a copy of a is assigned to b
	b[0] = "sh" // a中第一个元素不能被修改，依旧是'67'
	fmt.Println("a is ", a)
	fmt.Println("b is ", b) // b中第一个元素已被修改

	// -------- output
	//a is  [67 jinghua meihua 233 267]
	//b is  [sh jinghua meihua 233 267]

	// 数组的大小
	c := [3]int{5, 78, 8}
	var d [5]int
	//d = c //not possible since [3]int and [5]int are distinct types
	fmt.Println(reflect.TypeOf(c))
	fmt.Println(reflect.TypeOf(d))

	//     -------     output
	//[3]int
	//[5]int

}
