package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
)

// 给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。
// https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/
// 滑动窗口
func main() {
	res := lengthOfLongestSubstringCode("LiuQi")
	fmt.Println(res)
}

func lengthOfLongestSubstringCode(s string) int {
	// hash map，存储
	m := map[byte]int{}
	sLength := len(s)
	rk, ans := -1, 0
	for i := 0; i < sLength; i++ {
		if i != 0 {
			// 每次计数完成后，需要将上一轮已经计算过的字符从map中删除，该key为字节数组的s[i-1]，以该字符为首的计算已经完成了历史使命，需要删除掉，避免脏数据
			delete(m, s[i-1])
			logrus.Errorf("i ---> %d |删除元素为 ---- > %v | 删除后 -----> %v", i, s[i-1], m)
		}
		//rk最大为5，第二轮for的2个条件
		//1. rk +1 = sLength，最大只能向右移动5次
		//2. map类型的m[s[rk+1]]，当map中key=0，可以继续移动，因为key的value为0时，代表该key从没被set过=1，可以继续循环
		for rk+1 < sLength && m[s[rk+1]] == 0 {
			//m[s[rk+1]]++的意思是字节map中，s字节数组中，可以从0下标开始set该map的key为1，这也是rk为什么初始值设置为-1的原因，此过程中rk开始计数，rk的值和最大子串长度有关
			m[s[rk+1]]++
			rk++
			logrus.Info("rk----", rk)
		}
		// rk-i+1，减去i是因为多for了几轮，需要减去这个计数，+1是因为计数值是从-1开始的，需要+1补上
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
