package main

import "fmt"

// https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/x2cv1c/
func main() {
	res := plusOne([]int{9})
	fmt.Println(res)
}
func plusOne(digits []int) []int {

	// 计算出当前宿主中从最后一个元素开始，有多少个=9的元素，且不间断
	var num = 0
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] == 9 {
			num++
		} else {
			break
		}
	}
	// 没有9的情况，仅尾数+1
	if num == 0 {
		digits[len(digits)-1] = digits[len(digits)-1] + 1
	} else { //否则只需要将最后一个9前面的值+1,原来为9的值变为0就可以，需要得出最后一个连续的9在哪里
		// 如果9的个数==数组总长度，数组首位需要增加一个值为0的元素
		if len(digits) == num {
			digits = append([]int{0}, digits...)
		}
		digits[len(digits)-1-num] = digits[len(digits)-1-num] + 1
		for i := 0; i < num; i++ {
			digits[len(digits)-i-1] = 0
		}

	}
	//fmt.Println(num)
	////fmt.Println(digits)
	return digits
}
