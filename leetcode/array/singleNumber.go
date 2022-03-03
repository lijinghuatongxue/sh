package main

import (
	"fmt"
)

func main() {
	s := fmt.Sprintf("%d", singleNumberBitOperation([]int{1, 2, 6, 6, 1, 7, 7}))
	fmt.Println(s)
}
func singleNumberBitOperation(nums []int) int {
	//任何数和自己做异或运算，结果为 0; a^a=0
	//任何数和 0 做异或运算，结果还是自己; a^0 = a
	//异或运算中，满足交换律和结合律，也就是a^b^a = b^a^a = b^(a^a) = b^0 = b
	var res int
	for _, v := range nums {
		res = res ^ v
		fmt.Println(res)
	}
	return res
}
