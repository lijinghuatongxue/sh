package main

import (
	"fmt"
)

// 实现 strStr()
// https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xnr003/
func main() {
	fmt.Println(strStr("liuqi", "liu"))
}
func strStr(haystack string, needle string) int {
	// 全部为空的情况
	if len(haystack) == 0 && len(needle) == 0 {
		return 0
	}
	// "needle"长度大于"haystack"的长度
	if len(needle) > len(haystack) {
		return -1
	}
	// for 暴力
outer:
	for i := 0; i+len(needle) <= len(haystack); i++ {
		for j := range needle {
			if haystack[i+j] != needle[j] {
				continue outer
			}
		}
		return i
	}
	return -1
}

func strStrLeetCode(haystack string, needle string) int {
	n, m := len(haystack), len(needle)
outer:
	for i := 0; i+m <= n; i++ {
		for j := range needle {
			if haystack[i+j] != needle[j] {
				continue outer
			}
		}
		return i
	}
	return -1
}

// todo kmp算法
