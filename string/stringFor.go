package main

import (
	"fmt"
	"reflect"
)

func main()  {
	var s  = "AB"
	fmt.Println(s[0]," ",reflect.TypeOf(s[0]))
	for _, v := range s {
		fmt.Println(v," ",reflect.TypeOf(v))
	}
	x := []int{1,2,10,4,5,6}
	y := x[2:3]
	fmt.Println(y)

}
