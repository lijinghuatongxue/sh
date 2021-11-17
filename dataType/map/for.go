package main

import "fmt"

func main() {
	var mapInitDemo map[string]string
	mapInitDemo = make(map[string]string)
	mapInitDemo["key1"] = "value1"
	mapInitDemo["key2"] = "value2"
	// 遍历 map
	for key := range mapInitDemo {
		fmt.Println("Capital of", key, "is", mapInitDemo[key])
	}
}
