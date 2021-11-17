package main

import (
	"fmt"
)

// delete(map, key) 函数用于删除集合的元素, 参数为 map 和其对应的 key。删除函数不返回任何值。
func main() {
	var mapDemo map[string]string
	mapDemo = make(map[string]string)
	mapDemo = map[string]string{"C": "cValue", "A": "aValue", "B": "bValue"}
	fmt.Println("原始map ->", mapDemo)

	delete(mapDemo, "C")
	fmt.Println("删除操作后的map ->", mapDemo)

}
