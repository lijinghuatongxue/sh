package main

import "fmt"

type ListNodeReverseList struct {
	Val  int
	Next *ListNodeReverseList
}

// 给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。
// https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xnnhm6/
func main() {
	// 头链表
	var head = new(ListNodeReverseList)
	head.Val = 1
	fmt.Println("当前链表 head:", head)
	// 增加一些链表
	var node1 = new(ListNodeReverseList)
	node1.Val = 67
	var node2 = new(ListNodeReverseList)
	node2.Val = 6723
	head.Next = node1
	node1.Next = node2
	fmt.Printf("原始的 -> \n")
	forShowNode(head)
	fmt.Printf("反转后的 -> \n")
	forShowNode(reverseList(head))
}
func reverseList(head *ListNodeReverseList) *ListNodeReverseList {
	var prev *ListNodeReverseList
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}
func forShowNode(n *ListNodeReverseList) {
	for n != nil {
		fmt.Println(*n)
		n = n.Next
	}
}
