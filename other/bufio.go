package main

import (
	"bufio"
	"fmt"
	"os"
)

// 获取终端输入的内容,bufio pkg
func main() {
	fmt.Println("输入：")
	// 获取input内容
	inputContent := bufio.NewReader(os.Stdin)
	s1, _ := inputContent.ReadString('\n')
	fmt.Println("输入内容为：", s1)
}
