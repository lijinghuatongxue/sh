package main

import (
	"fmt"
	"sort"
)

// https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xnmav1/
// 最长公共前缀
func main() {
	fmt.Println(longestCommonPrefix([]string{"liuqi", "liu", "li"}))
}
func longestCommonPrefix(strs []string) string {
	// 为空的判断
	if len(strs) == 0 {
		return ""
	}
	// 排序
	sort.Strings(strs)
	//fmt.Println(strs)
	// 以最小长度的0下标的元素，作为循环限制
	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			if strs[j][i] != strs[0][i] || i == len(strs[j]) {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}
