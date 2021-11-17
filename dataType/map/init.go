package main

import "fmt"

/* 1. 声明变量，默认 map 是 nil */
//var map_variable map[key_data_type]value_data_type

/* 2. 使用 make 函数 */
//map_variable = make(map[key_data_type]value_data_type)

/* 使用 直接赋值 */
// map_variable := map[string]float32 {"C":5, "Go":4.5, "Python":4.5, "C++":2 }

func main() {
	var mapInitDemo map[string]string
	mapInitDemo = make(map[string]string)
	mapInitDemo["key1"] = "value1"
	mapInitDemo["key2"] = "value2"

	fmt.Println(mapInitDemo)
}
