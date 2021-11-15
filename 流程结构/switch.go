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
	case language == "Python":
		fmt.Printf("良好\n")
	case language == "Rust":
		fmt.Printf("还行\n")
	default:
		fmt.Printf("未知\n")
	}
}

//$ go run switch.go Python
//良好
