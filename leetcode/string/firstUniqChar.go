package main

import "fmt"

// 字符串中的第一个唯一字符
// https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xn5z8r/
func main() {
	fmt.Println(firstUniqChar("liuqi"))
}
func firstUniqChar(s string) int {
	cnt := [26]int{}
	for _, ch := range s {
		cnt[ch-'a']++
	}
	for i, ch := range s {
		if cnt[ch-'a'] == 1 {
			return i
		}
	}
	return -1
}
