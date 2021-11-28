package main

import (
	"fmt"
	"strings"
)

// 给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。
// https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/
// todo 滑动窗口
// ----------------------------------------------------
// 复杂解法：
// 1. 初始化一个map来存储，主要结构是 map[L:1 Li:1 Liu:1 LiuQ:1]，key为子字符串，value随便 int
// 2. 初始化一个resInt int类型，用来存储最大的无重复字符的子字符串的长度
// 3. for 原始字符串，累加一个无重复字符的子字符串，累加过程需要判断是否已存在相同字符
// 4. return  resInt
// ----------------------------------------------------
func main() {
	res := lengthOfLongestSubstring("LiuQi")
	fmt.Println(res)
}

func lengthOfLongestSubstring(s string) int {
	var resInt int
	var res string
	for i := 0; i < len(s); i++ {
		if i == 0 {
			for i := 0; i < len(s); i++ {
				// check是否有重复的字符
				if strings.Contains(res, string(s[i])) {
					break
				} else {
					res += string(s[i])
					// 如果当前res长度大于临时值(当前最长子字符串)
					if len(res) > resInt {
						resInt = len(res)
					}
				}

			}
		}
		res := ""
		if i != 0 && i < len(s) {
			for _, v := range s[i:] {
				// check是否有重复的字符
				if strings.Contains(res, string(v)) {
					break
				} else {
					res += string(v)
					// 如果当前res长度大于临时值(当前最长子字符串)
					if len(res) > resInt {
						resInt = len(res)
					}
				}

			}
		}
	}
	return resInt
}
