package main

import (
	"fmt"
	"sh/v1/dataType/interface/demo1/demo"
)

type T struct {
	demo.Controller
}

func (t *T) Get() {
	//new(test.Controller).Get()
	fmt.Print("T")
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
