package main

import (
	"fmt"
	"regexp"
	"strings"
)

// 给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。
// https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xne8id/
func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
}
func isPalindrome(s string) bool {
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		fmt.Println(err)
	}
	s = reg.ReplaceAllString(s, "")
	// 变成小写
	s = strings.ToLower(s)
	// 左右指针
	for left, right := 0, len(s)-1; left < len(s)/2; left++ {
		if s[left] != s[right] {
			return false
		}
		right--
	}
	return true
}
