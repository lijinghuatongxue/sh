package main

import "fmt"

type Num interface {
	Add(a int) int
	Subtract(b int) int
}
type NumDemo struct {
	SomeSum int
}

type NumDemo1 struct {
	SomeSum int
}

func (numAdd NumDemo) Subtract(someNum int) int {
	return someNum - 1
}

func (numAdd NumDemo) Add(someNum int) int {
	return someNum + 1
}

func (numAdd NumDemo1) Subtract(b int) int {
	return b - 1
}
func (numAdd NumDemo1) Add(someNum int) int {
	return someNum + 1
}

func main() {
	var res Num
	res = NumDemo{1}
	fmt.Println(res.Add(1))
	fmt.Println(res.Subtract(1))

	// second demo

	var res1 Num
	res1 = NumDemo1{1}
	fmt.Println(res1.Subtract(1))

}
