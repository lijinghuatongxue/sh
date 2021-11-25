package main

import (
	"fmt"
	"math"
	"strconv"
)

// https://leetcode-cn.com/problems/add-two-numbers/

func main() {
	V := [...]int{2, 4, 9, 5, 6, 4, 9}
	var l1 = new(ListNode)
	l1.Val = V[0]
	var l2 = new(ListNode)
	l2.Val = V[1]
	var l3 = new(ListNode)
	l3.Val = V[2]
	l1.Next = l2
	l2.Next = l3
	// 第2条链表
	var L1 = new(ListNode)
	L1.Val = V[3]
	var L2 = new(ListNode)
	L2.Val = V[4]
	var L3 = new(ListNode)
	L3.Val = V[5]
	var L4 = new(ListNode)
	L4.Val = V[6]
	L1.Next = L2
	L2.Next = L3
	L3.Next = L4

	res := addTwoNumbers(l1, L1)
	fmt.Println(res)
	//Show(res)

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	var SumInit, maxLen, iNum int
	var l1Len, l2Len, compareValue int
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
	// 获取链表长度差值，末尾补 zero value
	compareValue = int(math.Abs(float64(l2Len - l1Len)))
	if l1Len > l2Len {
		for i := 0; i <= compareValue; i++ {
			var node = ListNode{Val: 0}
			(tail).Next = &node
			tail = &node
		}
		//tail = l2
	} else {
		//tail = l1
		for i := 0; i <= compareValue; i++ {
			var node = ListNode{Val: 0}
			(tail).Next = &node
			tail = &node
		}
		//tail = l1
	}
	Show(L1)
	Show(L2)
	for i := 0; i < maxLen; i++ {
		if SumInit == 10 {
			SumInit = L1.Val + L2.Val + 10
		} else {
			SumInit = L1.Val + L2.Val
		}
		fmt.Println("当前sum为", SumInit)
		if SumInit >= 10 {
			SumInit = 10
		} else {
			SumInit = 0
		}
		fmt.Printf("(%d + %d)/10 = %d \n", L1.Val, L2.Val, (L1.Val+L2.Val)%10)
		fmt.Printf("第%d轮value | SumInit ：%d\n", i, SumInit)
		fmt.Println("当前value", L1.Val)
		if L1.Next == nil {
			L1.Next.Val = 0
			L1.Next.Next = nil
		}
		L1 = L1.Next
		if L2.Next == nil {
			L2.Next.Val = 0
			L2.Next.Next = nil
		}
		L2 = L2.Next
	}

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
