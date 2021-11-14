package main

import "fmt"

func main() {
	var a = []int{1, 2, 3}
	// ==================================== append ====================================
	// 尾部追加
	a = append(a, 4) // 追加一个元素
	fmt.Println(a)
	a = append(a, 5, 6, 7) //追加多个元素
	fmt.Println(a)
	a = append(a, []int{8, 9, 10}...) // 追加一个切片
	fmt.Println(a)
	// 开头追加  （⚠️ 会导致内存的重新分配，已有的元素重新复制一次）
	a = append([]int{0}, a...) // 开头追加一个元素，将原来的切片作为追加内容
	fmt.Println(a)
	a = append([]int{-16, -15}, a...) //开头追加一个切片，将原来的切片作为追加内容
	fmt.Println(a)
	// 中间追加
	a = append(a[:2], append([]int{-10}, a[1:]...)...) // 在下标为2的元素前面增加一个元素，值为-10 【思路：原来的slice被切割为2部分，以被追加位置作为分隔，:n & n: "n:"的基础上追加元素】
	fmt.Println(a)
	a = append(a[:2], append([]int{-13, -12, -11}, a[2:]...)...) // 在下标为2的元素前面增加一个slice，值为-10 【思路：原来的slice被切割为2部分，以被追加位置作为分隔，:n & n: "n:"的基础上追加slice】
	fmt.Println(a)

	// 任意下标位置i后插入元素，高性能版，copy + append
	var i = 2
	a = append(a, 0)     //切片扩展一个空间(元素)，提前预留
	copy(a[i+1:], a[i:]) // a[i:] 向后移动一个位置
	a[i] = -14           // 赋值
	fmt.Println(a)
	// 任意下标位置i后插入一个切片，高性能版，copy + append
	var n = []int{-115, -114}
	a = append(a, n...)       //切片扩展足够的空间(切片)，提前预留
	copy(a[i+len(n):], a[i:]) // a[i:] 向后移动一个len(n)
	copy(a[i:], n)            // 赋值
	fmt.Println(a)
}
