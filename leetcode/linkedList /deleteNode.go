package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 删除链表中的节点
// https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xnarn7/
func main() {
	var head = new(ListNode)
	deleteNode(head)
}

func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
