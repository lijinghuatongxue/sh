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
//
//
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
	//
	var existNum int
	for i := 0; i < len(s); i++ {
		if i == 0 {
			for i := 0; i < len(s); i++ {
				if strings.Contains(res, string(s[i])) {
					break
				} else {
					res += string(s[i])
					existNum = mapRes[res]
					if existNum == 0 {
						mapRes[res] = 1
					} else {
						mapRes[res] = existNum + 1
					}
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
					existNum = mapRes[res]
					if existNum == 0 {
						mapRes[res] = 1
					} else {
						mapRes[res] = existNum + 1
					}
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
