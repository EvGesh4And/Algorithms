package main

import "fmt"

type Node struct {
	val  int
	next *Node
}

type MyLinkedList struct {
	head *Node
}

func Constructor() MyLinkedList {
	return MyLinkedList{head: &Node{val: 0, next: nil}}
}

func skip(head *Node, index int) (*Node, bool) {

	current := head
	idx := -1

	for idx < index {
		current = current.next
		if current != nil {
			idx++
		} else {
			return current, false
		}
	}
	return current, true
}

func (this *MyLinkedList) Get(index int) int {
	if index < 0 {
		return -1
	}

	current := this.head

	current, ok := skip(current, index)

	if !ok {
		return -1
	}

	return current.val
}

func (this *MyLinkedList) AddAtHead(val int) {
	this.head.next = &Node{val: val, next: this.head.next}
}

func (this *MyLinkedList) AddAtTail(val int) {
	current := this.head

	for current.next != nil {
		current = current.next
	}

	current.next = &Node{val: val, next: nil}
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index < 0 {
		return
	}

	current := this.head
	idx := -1

	for idx < index-1 {
		current = current.next
		if current != nil {
			idx++
		} else {
			return
		}
	}

	current.next = &Node{val: val, next: current.next}
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 {
		return
	}

	current := this.head
	idx := -1

	for idx < index-1 {
		current = current.next
		if current != nil {
			idx++
		} else {
			return
		}
	}

	if current.next != nil {
		current.next = current.next.next
	}
}

func main() {
	m := Constructor()

	m.AddAtHead(1)
	m.AddAtTail(3)

	m.AddAtIndex(1, 2)

	m.DeleteAtIndex(1)

	fmt.Println(m.Get(0), m.Get(1))
}
