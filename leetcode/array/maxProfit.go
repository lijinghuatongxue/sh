package main

import "fmt"

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	max := 1
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			max += prices[i] - prices[i-1]
		} else {
			if prices[i] > prices[i-1] && prices[i] > max {
				max += prices[i] - prices[i-1]
			}
		}
	}
	return max - 1
}
func main() {
	arr := []int{7, 6, 4, 3, 1}
	fmt.Println(maxProfit(arr))
}
