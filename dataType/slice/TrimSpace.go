package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
)

// TrimSpace 消除byte切片中的空格元素
func TrimSpace(s []byte) []byte {
	b := s[:0]
	for _, i := range s {
		logrus.Info(i)
		// 32 = ' ' // https://baike.baidu.com/item/ASCII/309296
		if i != 32 {
			b = append(b, i)
		}
		//if i != ' '{
		//	b = append(b,i)
		//}
	}
	return b
}
func main() {
	s := []byte{'1', 1, 2, ' ', 5}
	result := TrimSpace(s)
	fmt.Println(result)
}
