package main

import "fmt"

// type 定义一个函数类型（如果函数的类型比较复杂，我们可以使用type来定义这个函数的类型）
type funcDemo func(a int, b int) int

func funcDemo1() funcDemo {
	funcReturn := func(a, b int) int {
		var res int
		res = a + b
		return res
	}
	return funcReturn
}
func main() {
	res := funcDemo1()
	fmt.Println(res(1, 2))
}
