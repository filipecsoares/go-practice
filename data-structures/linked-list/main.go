package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head   *Node
	length int
}

func (l *LinkedList) prepend(n *Node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}

func main() {
	list := LinkedList{}
	node1 := &Node{data: 20}
	node2 := &Node{data: 30}
	node3 := &Node{data: 15}
	list.prepend(node1)
	list.prepend(node2)
	list.prepend(node3)
	fmt.Println(list)
}
