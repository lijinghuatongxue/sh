package main

import "fmt"

// https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xnhbqj/
func main() {
	s := "liuqi" // iquil
	//s := "liuuqi" // iquuilß
	reverseString([]byte(s))
	fmt.Println(s)
}
func reverseString(s []byte) {
	for left, right := 0, len(s)-1; left < right; left++ {
		s[left], s[right] = s[right], s[left]
		right--
	}
}
