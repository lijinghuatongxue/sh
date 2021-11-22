package main

import "fmt"

type Node struct {
	data int   //  元素数据
	next *Node // 下一个节点
}

func showNode(p *Node) { //遍历
	for p != nil {
		fmt.Println(*p)
		p = p.next //移动指针
	}
}

func main() {
	var head = new(Node)
	head.data = 1
	fmt.Println("当前链表 head:", head)
	var node1 = new(Node)
	node1.data = 2
	fmt.Println("当前链表 node1:", node1)
	// 增加下一个节点
	head.next = node1
	fmt.Println("增加完节点后的当前链表 head:", head)
	var node2 = new(Node)
	node2.data = 3

	node1.next = node2
	showNode(head)
}
