package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"math"
)

// https://leetcode-cn.com/problems/add-two-numbers/

func main() {
	V := [...]int{9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9}
	var l1 = new(ListNode)
	l1.Val = V[0]
	var l2 = new(ListNode)
	l2.Val = V[1]
	var l3 = new(ListNode)
	var l4 = new(ListNode)
	var l5 = new(ListNode)
	var l6 = new(ListNode)
	var l7 = new(ListNode)
	l3.Val = V[2]
	l4.Val = V[1]
	l5.Val = V[1]
	l6.Val = V[1]
	l7.Val = V[1]
	l1.Next = l2
	l2.Next = l3
	l3.Next = l4
	l4.Next = l5
	l5.Next = l6
	l6.Next = l7
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
	Show(res)

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var l1Len, LastSum, l2Len, MaxLength, compareValue int
	var initLinkNode *ListNode
	//EndL2 = l2
	//EndL1 = l1
	//	制造2个数组，追加，and 获取两个链表的长度
	var l1Arr, l2Arr []int
	for l1 != nil {
		l1Arr = append(l1Arr, l1.Val)
		l1 = l1.Next //移动指针到下一个链表
		l1Len++
	}

	for l2 != nil {
		l2Arr = append(l2Arr, l2.Val)
		l2 = l2.Next //移动指针到下一个链表
		l2Len++
	}
	// 获取链表长度差值，数组末尾补 zero value
	compareValue = int(math.Abs(float64(l2Len - l1Len)))
	fmt.Println("差值为", compareValue)
	if len(l1Arr) >= len(l2Arr) {
		for i := 0; i < compareValue; i++ {
			l2Arr = append(l2Arr, 0)
		}
		//// 倒序
		//for i, j := 0, len(l2Arr)-1; i < j; i, j = i+1, j-1 {
		//	l2Arr[i], l2Arr[j] = l2Arr[j], l2Arr[i]
		//}
	} else {
		for i := 0; i < compareValue; i++ {
			l1Arr = append(l1Arr, 0)
		}
		//// 倒序
		//for i, j := 0, len(l1Arr)-1; i < j; i, j = i+1, j-1 {
		//	l1Arr[i], l1Arr[j] = l1Arr[j], l1Arr[i]
		//}
	}
	var res []int
	if l2Len > l1Len {
		MaxLength = l2Len
	} else {
		MaxLength = l1Len
	}
	logrus.Info(MaxLength)
	logrus.Warnf("1: %v \n 1: %v", l1Arr, l2Arr)
	for i := 0; i < MaxLength; i++ {
		if LastSum >= 10 {
			if len(l1Arr) != 0 && len(l2Arr) != 0 {
				LastSum = l1Arr[i] + l2Arr[i] + 1
			}

		} else {
			if len(l1Arr) != 0 && len(l2Arr) != 0 {
				LastSum = l1Arr[i] + l2Arr[i]
			}
		}
		if LastSum >= 10 {
			//fmt.Println("当前sum为", LastSum)
			res = append(res, LastSum-10)
		} else {
			//fmt.Println("当前sum为", LastSum)
			res = append(res, LastSum)
		}
		if len(l1Arr) != 0 && len(l2Arr) != 0 {
			fmt.Printf(" 第%d轮value | SumInit ：%d |(%d + %d)/10 = %v \n", i, LastSum, l1Arr[i], l2Arr[i], res)
		}
		//if  EndL1 != nil || EndL2 != nil{
		//	EndL1 = EndL1.Next
		//	EndL2 = EndL2.Next
		//}

		// 判断最后一位，>= 10 ; 补 int 1
		if i+1 == len(l1Arr) {
			if LastSum >= 10 {
				res = append(res, 1)
			}
		}
	}
	fmt.Println("end -------> ", res)
	// 倒序
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	for _, i := range res {
		var node = ListNode{Val: i} // 赋值给 tail 链表
		//initLinkNode.Next= &node           //将新插入的node的next指向头结点
		//initLinkNode = &node                   //重新赋值头结点
		node.Next = initLinkNode
		initLinkNode = &node
	}
	return initLinkNode
}

func Show(p *ListNode) { //遍历
	for p != nil {
		fmt.Println("遍历 elementInt ：", p.Val)
		fmt.Println("遍历 链表 ：", *p)
		p = p.Next //移动指针到下一个链表
	}
}
