package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 1}
	fmt.Println(RemoveDuplicates(nums))
}
func RemoveDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	slow := 1
	for fast := 1; fast < len(nums); fast++ {
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}
