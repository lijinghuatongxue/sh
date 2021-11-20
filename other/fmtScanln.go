package main

import "fmt"

func main() {
	var a int
	var b string
	fmt.Println("请输入 [年龄 姓名]:")
	_, err := fmt.Scanf("%d %s", &a, &b)
	if err != nil {
		return
	}
	fmt.Printf("年龄: %d ｜姓名：%s", a, b)
}
