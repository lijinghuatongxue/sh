package main

import "fmt"

// https://github.com/lijinghuatongxue/sh/blob/main/%E6%B5%81%E7%A8%8B%E7%BB%93%E6%9E%84/switch/switch-check-interface-value-type.go
// 断言其实还有另一种形式，就是用在利用 switch语句判断接口的类型。每一个case会被顺序地考虑。当命中一个case 时，就会执行 case 中的语句，因此 case 语句的顺序是很重要的，因为很有可能会有多个 case匹配的情况。
func main() {
	var student interface{} = new(Student)
	s, ok := student.(Student) //安全，断言失败，也不会panic，只是ok的值为false
	if ok {
		fmt.Println(s.Name)
	}
}

type Student struct {
	Name string
}
