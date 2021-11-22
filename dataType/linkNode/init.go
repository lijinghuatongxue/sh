package main

import "fmt"

//  定义数据类型
type element string

// 定义节点
type linkNode struct {
	Data element   // 元素数据
	Next *linkNode // 指向下一个节点
}

// HeadLinkNode 链表头
type HeadLinkNode struct {
	length int
	Node   *linkNode
}

// LinkNoder 定义节点接口
type LinkNoder interface {
	Add(node element)               // 增加尾部
	Remove(index int) error         // 删除指定位置的node
	Insert(index int, node element) // 指定位置插入节点
	Len() int
	Search() int
	Get(index int) *linkNode
}

// New 初始化
func New() *HeadLinkNode {
	return &HeadLinkNode{length: 0, Node: &linkNode{Data: "67", Next: nil}}
}

// Add 增加末尾节点
func (h *HeadLinkNode) Add(node element) {
	l := h.Node
	for {
		if l.Next == nil {
			newNode := &linkNode{Data: node, Next: nil}
			l.Next = newNode
			break
		} else {
			l = l.Next
		}
	}
	h.length++
}

func main() {
	l := New()
	fmt.Println(l)
	//l.Add(element("liuqi"))
	fmt.Println(l.Node.Data)
}
