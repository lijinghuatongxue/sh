package main

import (
	"fmt"
	"sort"
	"strings"
)

// 给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。
// https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/
// todo 滑动窗口
// ----------------------------------------------------
// 复杂解法：
// 1. 初始化一个map来存储，主要结构是 map[L:1 Li:1 Liu:1 LiuQ:1]，key为子字符串，value随便 int
// 2. 获取所有子字符串，存取到map之前判断是否当前str已包含，已包含则break
// 3. 将当前map中的key全部追加到一个array
// 4. array进行顺序排序，获取最后一个值的length去return
// ----------------------------------------------------
func main() {
	res := lengthOfLongestSubstring("LiuQi")
	fmt.Println(res)
}

func lengthOfLongestSubstring(s string) int {
	var resArrInt []int
	var res string
	// 初始化一个map
	var mapRes map[string]int
	mapRes = make(map[string]int)
	for i := 0; i < len(s); i++ {
		if i == 0 {
			for i := 0; i < len(s); i++ {
				if strings.Contains(res, string(s[i])) {
					break
				} else {
					res += string(s[i])
					mapRes[res] = 1
				}

			}
		}
		res := ""
		if i != 0 && i < len(s) {
			for _, v := range s[i:] {
				if strings.Contains(res, string(v)) {
					break
				} else {
					res += string(v)
					mapRes[res] = 1
				}

			}
		}
	}
	fmt.Println(mapRes)
	for k := range mapRes {
		resArrInt = append(resArrInt, len(k))
	}
	sort.Ints(resArrInt)
	if len(resArrInt) != 0 {
		return resArrInt[len(resArrInt)-1]
	}
	return 0
}
