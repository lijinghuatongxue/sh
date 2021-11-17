package main

import "fmt"

func main() {
	// 与切片相似，映射是引用类型。当将映射分配给一个新变量时，它们都指向相同的内部数据结构。因此，一个的变化会反映另一个。
	var initMapDemo map[string]string
	initMapDemo = make(map[string]string)
	initMapDemo = map[string]string{"key1": "value1", "key2": "value2"}
	fmt.Println("原始map ->", initMapDemo)
	NewInitMapDemo := initMapDemo
	NewInitMapDemo["key2"] = "value3"
	fmt.Println("修改后的变量map | 未改变", NewInitMapDemo)
	fmt.Println("分配后的原始map | 跟随变量map元素改变", initMapDemo)
	// map value 比较
	if NewInitMapDemo["key1"] == initMapDemo["key1"] {
		fmt.Println("value is Equal")
	}
	// map 比较，不能比较
	//if NewInitMapDemo == initMapDemo{
	//
	//}
}
