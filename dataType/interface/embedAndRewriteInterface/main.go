package main

import (
	"fmt"
	"sh/v1/dataType/interface/embedAndRewriteInterface/demo"
)

type T struct {
	demo.Controller
}

func (t *T) Get() {
	//new(test.Controller).Get()
	fmt.Print("重写demo中的Get方法")
}

func (t *T) Post() {
	fmt.Print("T")
}

func main() {
	var something demo.Something
	something = new(T)
	var t T
	t.M = 1
	something.Get()
}
