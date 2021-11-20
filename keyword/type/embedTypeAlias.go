package main

import "fmt"

type people struct {
	name string
}

func (p people) show() string {
	fmt.Println("people name is", p.name)
	return p.name
}

// 别名 结构体
type person people

// student 嵌入两个结构
type student struct {
	people
	person
}

func (p person) show2() string {
	fmt.Println("person name is", p.name)
	return p.name
}

// type 别名结构体，新结构体值改变，旧结构体值不跟随改变
func main() {
	var s student
	// people name is 67
	s.people.name = "67"
	// person name is ad
	s.person.name = "ad"

	s.person.show2()
	s.people.show()

}
