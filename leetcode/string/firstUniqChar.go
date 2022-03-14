package main

import "fmt"

// 给定一个字符串 s ，找到 它的第一个不重复的字符，并返回它的索引 。如果不存在，则返回 -1 。
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
