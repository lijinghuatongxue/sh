package main

import (
	"fmt"
)

func main() {
	s := fmt.Sprintf("%d", singleNumber([]int{1, 2, 6, 6, 1, 7, 7}))
	fmt.Println(s)
}
func singleNumber(nums []int) int {
	var res int
	for _, v := range nums {
		res = res ^ v
		fmt.Println(res)
	}
	return res
}
