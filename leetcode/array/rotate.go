package main

import "fmt"

func rotate(nums []int, k int) {
	newNums := make([]int, len(nums))
	for i, v := range nums {
		newNums[(i+k)%len(nums)] = v
	}
	copy(nums, newNums)
}
func main() {
	rotate([]int{1, 2, 3, 4, 5, 6}, 2)
	//           0  1  2  3  4  5
	fmt.Println(3 % 6)
}
