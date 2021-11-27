package main

import (
	"fmt"
)

func main() {
	var res string
	var s = []string{"l", "i", "j", "i", "n", "g"}
	var resArr []string
	for i := 0; i < len(s); i++ {
		if i == 0 {
			for i := 0; i < len(s); i++ {
				fmt.Printf("当前index: %d ,value:%s\n", i, s[i])
				res += s[i]
				fmt.Println("当前拼接结果", res)
				resArr = append(resArr, res)
			}
		}
		res := ""
		if i != 0 && i < len(s) {
			for _, v := range s[i:] {
				fmt.Printf("当前index: %d ,value:%s\n", i, v)
				res += v
				fmt.Println("当前拼接结果", res)
				resArr = append(resArr, res)
			}
		}
		fmt.Println("----", i)
		fmt.Println("结果", res)
		fmt.Println("总结果", resArr)
	}
	//fmt.Println(s[1:])
}
