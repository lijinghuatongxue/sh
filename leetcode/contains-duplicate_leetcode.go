package main

import "log"

// https://leetcode-cn.com/problems/contains-duplicate/solution/cun-zai-zhong-fu-yuan-su-by-leetcode-sol-iedd/
// hash  table  解法
func containsDuplicateLeetcode(nums []int) bool {
	set := map[int]struct{}{}
	for _, i := range nums {
		if _, has := set[i]; has {
			return true
		}
		set[i] = struct{}{}
	}
	return false
}
func main() {
	var a = []int{1, 2, 3, 2, 3}
	res := containsDuplicateLeetcode(a)
	log.Println(res)
}
