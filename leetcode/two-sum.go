package main

import "fmt"

// todo hashmap
//

func twoSumSecond(nums []int, target int) []int {
	// init a map
	numMap := map[int]int{}
	for i, k := range nums {
		if p, ok := numMap[target-k]; ok {
			return []int{p, i}
		}
		numMap[k] = i
	}
	return nil
}

// https://leetcode-cn.com/problems/two-sum/

// 穷举
func twoSum(nums []int, target int) []int {
	var resArr []int
	var i int
	for i = 0; i < len(nums); i++ {
		var s int
		for s = 0; s < len(nums); s++ {
			if s == i {
				continue
			} else {
				if nums[i]+nums[s] == target {
					resArr := append(resArr, i)
					resArr = append(resArr, s)
					return resArr
				}
			}
		}
	}
	return resArr
}
func main() {
	inputArray := []int{1, 3, 4, 5, 9}
	var target = 10
	fmt.Println(twoSum(inputArray, target))
	fmt.Println(twoSumSecond(inputArray, target))
}
