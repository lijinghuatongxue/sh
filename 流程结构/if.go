package main

import (
	"fmt"
)

// https://github.com/rubyhan1314/Golang-100-Days/blob/master/Day01-15(Go%E8%AF%AD%E8%A8%80%E5%9F%BA%E7%A1%80)/day04_%E5%88%86%E6%94%AF%E8%AF%AD%E5%8F%A5.md#22-if-%E5%8F%98%E4%BD%93
func main() {
	if num := 10; num%2 == 0 { // 判断是否是偶数
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}
}
