package helpers

import (
	"fmt"
)

type Node struct {
	prev *Node
	next *Node
	key  interface{}
}

type LinkedList struct {
	head *Node
	tail *Node
}

func (L *LinkedList) Insert(key interface{}) {
	list := &Node{
		next: L.head,
		key:  key,
	}
	if L.head != nil {
		L.head.prev = list
	}
	L.head = list

	l := L.head
	for l.next != nil {
		l = l.next
	}
	L.tail = l
}

func (N *Node) StepForward() {
	N = N.next
}

func (l *LinkedList) Display() {
	list := l.head
	for list != nil {
		fmt.Printf("%+v ->", list.key)
		list = list.next
	}
	fmt.Println()
}

func (l *LinkedList) DisplayString() (returnString string) {
	list := l.head
	for list != nil {
		returnString = returnString + fmt.Sprintf("%+v ->", list.key)
		list = list.next
	}
	return returnString
}

func Display(list *Node) {
	for list != nil {
		fmt.Printf("%v ->", list.key)
		list = list.next
	}
	fmt.Println()
}

func ShowBackwards(list *Node) {
	for list != nil {
		fmt.Printf("%v <-", list.key)
		list = list.prev
	}
	fmt.Println()
}

func (l *LinkedList) Reverse() {
	curr := l.head
	var prev *Node
	l.tail = l.head

	for curr != nil {
		next := curr.next
		curr.next = prev
		prev = curr
		curr = next
	}
	l.head = prev
	Display(l.head)
}

func (l *LinkedList) Length() (length int) {
	count := 0
	list := l.head
	if list == nil {
		return count
	} else {
		count = 1
		for list.next != nil {
			count = count + 1
			list = list.next
		}
	}

	return count
}

func (l *LinkedList) GetHead() (node Node) {
	return *l.head
}

func (l *LinkedList) SetHead(node Node) {
	*l.head = node
}

func (l *LinkedList) GetTail() (node Node) {
	return *l.tail
}

func (l *LinkedList) SetTail(node Node) {
	*l.tail = node
}

// get the key for int type
func (n *Node) GetKeyInt() (key int) {
	return n.key.(int)
}

// set the key for int type
func (n *Node) SetKeyInt(input int) {
	n.key = input
}
