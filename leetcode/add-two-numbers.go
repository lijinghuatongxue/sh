package main

import (
	"fmt"
	"math"
	"strconv"
)

// https://leetcode-cn.com/problems/add-two-numbers/

func main() {
	V := [...]int{2, 4, 3, 5, 6, 4}
	var l1 = new(ListNode)
	l1.Val = V[0]
	var l2 = new(ListNode)
	l2.Val = V[1]
	var l3 = new(ListNode)
	l3.Val = 3
	l1.Next = l2
	l2.Next = l3
	// 第2条链表
	var L1 = new(ListNode)
	L1.Val = V[3]
	var L2 = new(ListNode)
	L2.Val = V[4]
	var L3 = new(ListNode)
	L3.Val = V[5]
	L1.Next = L2
	L2.Next = L3

	res := addTwoNumbers(l1, L1)
	fmt.Println(res)
	Show(res)

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	var i, L, iNum int
	var l1Len, l2Len int
	var L1, L2, tail *ListNode
	L1 = l1
	L2 = l2
	//tail = ListNode
	//fmt.Println("tail",tail)
	//	获取两个链表的长度
	for l1 != nil {
		l1 = l1.Next //移动指针到下一个链表
		l1Len++
	}

	for l2 != nil {
		l2 = l2.Next //移动指针到下一个链表
		l2Len++
	}
	//fmt.Println("len",l2Len,l1Len)

	for L1 != nil {
		//fmt.Println(L1.Val,"*",int(math.Pow10(l1Len)))
		iNum += L1.Val * (int(math.Pow10(l1Len)))
		L1 = L1.Next //移动指针到下一个链表
		i++
		l1Len--
	}
	for L2 != nil {
		//fmt.Println(L2.Val,"*",int(math.Pow10(l2Len)))
		iNum += L2.Val * (int(math.Pow10(l2Len)))
		L2 = L2.Next //移动指针到下一个链表
		L++
		l2Len--
	}
	//fmt.Println("分别求出两个链表携带的val 之和 sum:",iNum/10)

	// 将sum拆开塞入到一个新的链表中

	//tail = l1
	var slice []int
	for _, i := range fmt.Sprintf("%d", iNum/10) {
		//fmt.Println(i)
		var a int
		a, _ = strconv.Atoi(string(i))
		//fmt.Println(a)
		slice = append(slice, a)
	}
	fmt.Println(slice)
	for _, i := range slice {
		var node = ListNode{Val: i} // 赋值给 tail 链表
		//(*tail).Next= &node           //将新插入的node的next指向头结点
		//tail = &node                   //重新赋值头结点
		node.Next = tail
		tail = &node
	}
	return tail
}

func Show(p *ListNode) { //遍历
	for p != nil {
		fmt.Println("遍历 elementInt ：", p.Val)
		fmt.Println("遍历 链表 ：", *p)
		p = p.Next //移动指针到下一个链表
	}
}
