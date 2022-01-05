package main

import (
	"log"
	"sort"
)

// https://leetcode-cn.com/problems/contains-duplicate/submissions/
// 先排序，判断除最后一个元素之外的所有元素是否等于后一个元素，排除最后一个元素的原因是，防止数组越界
func containsDuplicate(nums []int) bool {
	sort.Ints(nums)
	for i := range nums {
		if i != len(nums)-1 {
			if nums[i] == nums[i+1] {
				return true
			}
		}
	}
	return false
}
func main() {
	var a = []int{1, 2, 3, 2, 3}
	res := containsDuplicate(a)
	log.Println(res)
}
