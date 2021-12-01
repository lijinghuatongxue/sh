package main

import (
	"fmt"
	"sort"
)

// https://leetcode-cn.com/problems/median-of-two-sorted-arrays/
// 二分查找
// todo

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	nums1 = append(nums1, nums2...)
	sort.Ints(nums1)
	fmt.Println(nums1)
	var res float64
	if len(nums1)%2 == 0 {
		res = float64(nums1[(len(nums1)/2)-1]+nums1[len(nums1)/2]) / 2
	} else {
		res = float64(nums1[(len(nums1) / 2)])
	}
	return res
}
func main() {
	var nums1 = []int{1, 8}
	var nums2 = []int{3, 4}
	res := findMedianSortedArrays(nums1, nums2)
	fmt.Println(res)

}
