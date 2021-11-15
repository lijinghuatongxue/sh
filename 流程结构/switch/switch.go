package main

import (
	"fmt"
	"os"
)

func main() {
	var language string
	language = os.Args[1]
	switch {
	case language == "Golang":
		fmt.Printf("优秀!\n")
		fallthrough //如需贯通后续的一个case，就添加fallthrough
	case language == "Python":
		fmt.Printf("良好\n")
	case language == "Rust":
		fmt.Printf("还行\n")
	default:
		fmt.Printf("未知\n")
	}
}

//$ go run switch.go Golang
//优秀!
//良好
