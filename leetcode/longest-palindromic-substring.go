package main

import (
	"github.com/sirupsen/logrus"
)

func main() {
	//    012345678910
	s := "aaaakaaaajjakajj"
	res := longestPalindrome(s)
	logrus.Println(res)
}
func longestPalindrome(s string) string {
	var spaceNum, strNum int
	var spaceMap map[int]int
	spaceMap = make(map[int]int)
	var strMap map[int]int
	strMap = make(map[int]int)
	if len(s) == 1 {
		return s
	}
	if len(s) == 0 {
		return s
	}
	//情况1. wuilllqqqi，空格情况作为中间间隔
	for i := 1; i < len(s); i++ {
		if i == len(s)-1 {
			break
		}
		forNum := Min(i, len(s)-i-1)
		for n := 0; n < forNum; n++ {
			//logrus.Errorf("%d | %d |i -> %d |n -> %d",s[i],s[i+1+n],i,n)
			if s[i-n] == s[i+1+n] {
				//logrus.Info(n+1)
				spaceNum = Max(spaceNum, n+1)
				if spaceNum > spaceMap[spaceNum] {
					spaceMap[spaceNum] = i
				}
			}
		}
	}
	//情况2. wuillliqqqi，字符情况作为中间间隔
	for i := 0; i < len(s); i++ {
		if i == len(s)-1 {
			break
		}
		forNum := Min(i, len(s)-i)
		for n := 1; n < forNum; n++ {
			if s[i-n] == s[i+n] {
				logrus.Errorf("%d | %d |i -> %d |n -> %d", s[i-n], s[i+n], i, n)
				if n > strNum {
					strMap[n] = i
				}
				strNum = Max(strNum, n)
			}
		}

	}
	logrus.Warnf("最大value: %d", spaceNum)
	logrus.Warnf("map strMap: %v", strMap)

	logrus.Warnf("最大value: %d", strNum)
	logrus.Warnf("map spaceMap: %v", spaceMap)

	if spaceNum > strNum {
		return s[spaceMap[spaceNum]-spaceNum+1 : spaceMap[spaceNum]+spaceNum+1]
	} else {
		return s[strMap[strNum]-strNum+1 : strMap[strNum]+strNum]
	}
}
func Max(a, b int) int {
	//logrus.Warnf("对比Max a: %d |b: %d", a, b)
	if a > b {
		return a
	}
	return b
}
func Min(a, b int) int {
	//logrus.Infof("对比Min a: %d |b: %d", a, b)
	if a > b {
		return b
	}
	return a
}
