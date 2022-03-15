package main

import (
	"fmt"
	"sort"
)

// 给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词
// https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xn96us/
func main() {
	fmt.Println(isAnagram("liuqi", "iquli"))
}
func isAnagram(s string, t string) bool {
	// 先判断长度是否一致，第一次判断
	if len(s) != len(t) {
		return false
	}
	// 转化为数组
	sArr := []byte(s)
	tArr := []byte(t)
	sort.Slice(sArr, func(i, j int) bool {
		return sArr[i] < sArr[j]
	})
	sort.Slice(tArr, func(i, j int) bool {
		return tArr[i] < tArr[j]
	})
	return string(sArr) == string(tArr)
}
