package main

import (
	"fmt"
)

// 给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。
// https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/
// 滑动窗口
func main() {
	res := lengthOfLongestSubstringCode("abcabb")
	fmt.Println(res)
}

func lengthOfLongestSubstringCode(s string) int {
	m := map[byte]int{}
	sLength := len(s)
	rk, ans := -1, 0
	for i := 0; i < sLength; i++ {
		if i != 0 {
			delete(m, s[i-1])
		}
		for rk+1 < sLength && m[s[rk+1]] == 0 {
			m[s[rk+1]]++
			rk++
		}
		ans = max(ans, rk-i+1)
	}
	return ans
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
