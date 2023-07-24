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

func (l LinkedList) print() {
	for n := l.head; n != nil; n = n.next {
		fmt.Printf("%d ", n.data)
	}
	fmt.Println()
}

func (l *LinkedList) deleteByValue(value int) *Node {
	if l.length == 0 {
		return nil
	}
	if l.head.data == value {
		l.head = l.head.next
		l.length--
		return l.head
	}
	prev := l.head
	for prev.next.data != value {
		if prev.next.next == nil {
			return nil
		}
		prev = prev.next
	}
	prev.next = prev.next.next
	l.length--
	return prev.next
}

func main() {
	list := LinkedList{}
	node1 := &Node{data: 20}
	node2 := &Node{data: 30}
	node3 := &Node{data: 15}
	node4 := &Node{data: 25}
	node5 := &Node{data: 35}
	list.prepend(node1)
	list.prepend(node2)
	list.prepend(node3)
	list.prepend(node4)
	list.prepend(node5)
	list.print()
	fmt.Println(list.deleteByValue(15))
	list.print()
}
