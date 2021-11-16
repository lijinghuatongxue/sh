package main

import "fmt"

// 数组初始化姿势
func main() {
	var a [4]float32 // 等价于：var arr2 = [4]float32{}
	fmt.Println(a)   // [0 0 0 0]
	var b = [5]string{"ruby", "王二狗", "rose"}
	fmt.Println(b)                          // [ruby 王二狗 rose  ]
	var c = [5]int{'A', 'B', 'C', 'D', 'E'} // byte
	fmt.Println(c)                          // [65 66 67 68 69]
	d := [...]int{1, 2, 3, 4, 5}            // 根据元素的个数，设置数组的大小
	fmt.Println(d)                          //[1 2 3 4 5]
	e := [5]int{4: 100}                     // [0 0 0 0 100] ,4: 代表[:]前面四个元素 + 100
	fmt.Println(e)
	f := [...]int{0: 1, 4: 1, 9: 1} // 9: 代表已有9个元素，加上后面的1，共10个元素，[n: value]代表下标为n的元素为 value
	fmt.Println(f)                  // [1 0 0 0 1 0 0 0 0 1]

}
