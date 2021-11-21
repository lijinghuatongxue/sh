package main

import "fmt"

func main() {
	var a int
	a = 68      // 64 32 16 8 4 2 1 (64 + 4 = 68 | Binary -> 1000100)
	c := a & 67 // 64 + 2 + 1 = 67 | Binary -> 1000011
	// 1000100
	// 1000011
	// 1000000
	// 1   0  0  0  0 0 0
	// 64 32 16 8 4 2 1
	// 64
	fmt.Println(c) // 64
}
