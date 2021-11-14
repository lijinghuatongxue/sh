// init slice
package main

import "fmt"

// make出来的slice每个原色的值都是该数据类型的默认值
func main() {
	var a []int               // nil切片，和nil相等，一般用来表示一个不存在的切片
	var b = []int{}           // 空切片，和nil不相等，一般用来表示一个空的集合
	var c = []int{1, 2, 3}    // 有3个元素的切片，len和cap都为3
	var d = c[:2]             // 有2个元素的切片，len为2，cap为3
	var e = c[0:2:cap(c)]     // 有2个元素的切片，len为2，cap为3
	var f = c[:0]             // 有0个元素的切片，len为0，cap为3
	var g = make([]int, 3)    // 有3个元素的切片，len和cap都为3
	var h = make([]int, 2, 3) // 有2个元素的切片，len为2，cap为3
	var i = make([]int, 0, 3) // 有0个元素的切片，len为0，cap为3
	fmt.Printf("a :%v\nb :%v\nc :%v\nd :%v\ne :%v\nf :%v\ng :%v\nh: %v\ni :%v", a, b, c, d, e, f, g, h, i)
}
