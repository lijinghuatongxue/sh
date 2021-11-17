package main

import "fmt"

func main() {
	var mapInitDemo map[string]string
	mapInitDemo = make(map[string]string)
	mapInitDemo["key1"] = "value1"
	mapInitDemo["key2"] = "value2"
	// 根据key查找value
	/* 查看元素在集合中是否存在 */
	value, isOk := mapInitDemo["key1"]
	if isOk {
		fmt.Println("value is", value)
	}
}
