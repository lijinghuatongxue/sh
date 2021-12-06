package main

import (
	"github.com/sirupsen/logrus"
)

func main() {
	s := "bbbabbba"
	res := longestPalindrome(s)
	logrus.Println(res)
}
func longestPalindrome(s string) string {
	var resMap map[int]int
	resMap = make(map[int]int)
	var num, numTwo, numResTwo int
	logrus.Warnf("len /2 = %d", len(s)%2)
	//if len(s)%2 == 0 {
	//	for i := 0; i < len(s)/2; i++ {
	//		if s[i] == s[len(s)-1-i] {
	//			numThree++
	//		} else {
	//			break
	//		}
	//	}
	//}
	//logrus.Warn("最长子串不以0开头的字符长度 numThree 为", numThree)
	//if numThree*2 == len(s) {
	//	return s
	//}
	for i := 0; i < len(s)-1; i++ {
		if i == 0 {
			forNum := len(s) - 1
			for m := 1; m < forNum; m++ {
				if s[0] == s[m] {
					num = m
				} else {
					break
				}
			}
		} else {
			// 求出左滑右滑最小次数，避免越界
			forNum := Min(i, len(s)-i)
			for n := 1; n < forNum; n++ {
				if s[i-1] == s[i+1] {
					logrus.Errorf("%d ｜ %d |n %d ｜i %d", s[i-1], s[i+1], n+1, i)
					numTwo = n + 1
					resMap[numTwo] = i
				}
			}
			numResTwo = Max(numTwo, numResTwo)
		}
	}
	logrus.Info("最长子串以0开头的字符长度为", num)
	logrus.Warn("最长子串不以0开头的字符长度 numTwo 为", numResTwo)

	//var resArr []string
	//resArr =strings.Split(s,"")
	//if num > numTwo{
	//	return s[:num+1]
	//}else {
	//	//return s[map[3]]
	//}
	logrus.Info("map", resMap)
	return "1"
}
func Max(a, b int) int {
	logrus.Warnf("对比a: %d |b: %d", a, b)
	if a > b {
		return a
	}
	return b
}
func Min(a, b int) int {
	//logrus.Warnf("对比a: %d |b: %d", a, b)
	if a > b {
		return b
	}
	return a
}
