package main

import "fmt"

func main() {
	var a int
	fmt.Println(a)
	// int 别名为 myInt
	type myInt int
	//type myInt2 = int
	var b myInt
	fmt.Println(b)
}
