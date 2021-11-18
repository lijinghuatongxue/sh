package main

import "fmt"

type name int8
type first struct {
	a int
	b bool
	name
}

func main() {
	var a = first{1, false, 9}
	var b *first = &a
	fmt.Println(b.a, b.name, b.b)
	fmt.Println(a.a, a.name, a.b)
}
