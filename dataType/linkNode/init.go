package main

import "fmt"

type Node struct {
	elementInt int   //  元素数据
	next       *Node // 下一个节点
}

//             --------------   for linkNode   --------------
func showNode(p *Node) { //遍历
	for p != nil {
		fmt.Println("遍历 elementInt ：", p.elementInt)
		fmt.Println("遍历 链表 ：", *p)
		p = p.next //移动指针到下一个链表
	}
}

// 			   --------------   头部插入   --------------
func headInsert(p *Node) *Node {
	var tail *Node
	tail = p
	// 制造一个for去赋值 elementInt
	for i := 1; i < 10; i++ {
		var node = Node{elementInt: i} // 赋值给 tail 链表
		node.next = tail               //将新插入的node的next指向头结点
		tail = &node                   //重新赋值头结点
	}
	return tail
}

// 			   --------------   尾部插入   --------------
func tailInsert(p *Node) *Node {
	var tail *Node
	tail = p
	// 制造一个for去赋值 elementInt
	for i := 1; i < 10; i++ {
		var node = Node{elementInt: i} // 赋值给 tail 链表
		(*tail).next = &node           //将新插入的node的next指向头结点
		tail = &node                   //重新赋值头结点
	}
	return tail
}

func main() {
	// 头链表
	var head = new(Node)
	head.elementInt = 1
	fmt.Println("当前链表 head:", head)
	// 第二个链表
	var node1 = new(Node)
	node1.elementInt = 2
	fmt.Println("当前链表 node1:", node1)
	// 主链增加一个节点
	head.next = node1
	fmt.Println("增加完节点后的当前链表 head:", head)
	// 第三个链表
	var node2 = new(Node)
	node2.elementInt = 3
	// 第二个链表增加一个节点
	node1.next = node2
	// 头部插入
	res := headInsert(head)
	showNode(res)

	// 尾部插入
	res = tailInsert(head)
	showNode(res)

	// 循环
	showNode(head)

}
