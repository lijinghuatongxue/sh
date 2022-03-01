package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 1, 3, 4, 5, 5}
	fmt.Println(RemoveDuplicates(nums))
	fmt.Println(RemoveDuplicatesTwo(nums))
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

func RemoveDuplicatesTwo(nums []int) int {
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
