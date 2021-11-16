package main

import "fmt"

func main() {

	arr := [...]int{6, 7, 6, 7, 6, 7}
	// range
	for _, v := range arr {
		fmt.Printf("range |value -> %d\n", v)
	}
	fmt.Printf("\n")
	// index
	for i := 0; i < len(arr); i++ {
		fmt.Printf("index |value -> %d\n", arr[i])
	}
}
