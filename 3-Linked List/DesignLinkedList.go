package main

type Node struct {
	val  int
	next *Node
}

type MyLinkedList struct {
	head *Node
}

func Constructor() MyLinkedList {
	return MyLinkedList{}
}

func (this *MyLinkedList) Get(index int) int {
	if index < 0 || this.head == nil {
		return -1
	}

	idx := 0
	currentNode := this.head

	for idx != index && currentNode.next != nil {
		currentNode = currentNode.next
		idx++
	}

	if idx == index {
		return currentNode.val
	}

	return -1
}

func (this *MyLinkedList) AddAtHead(val int) {
	this.head = &Node{val: val, next: this.head}
}

func (this *MyLinkedList) AddAtTail(val int) {
	if this.head == nil {
		this.head = &Node{val: val}
		return
	}

	currentNode := this.head

	for currentNode.next != nil {
		currentNode = currentNode.next
	}

	currentNode.next = &Node{val: val}
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if this.head == nil && index > 0 {
		return
	}
	if index == 0 {
		if this.head == nil {
			this.AddAtHead(val)
		} else {
			newNode := &Node{val: val, next: this.head}
			this.head = newNode
		}
		return
	}

	idx := 0
	currentNode := this.head

	for idx != index-1 && currentNode.next != nil {
		currentNode = currentNode.next
		idx++
	}

	if idx == index-1 {
		newNode := &Node{val: val, next: currentNode.next}
		currentNode.next = newNode
	}

}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || this.head == nil {
		return
	}

	if index == 0 {
		this.head = this.head.next
		return
	}

	idx := 0
	currentNode := this.head

	for idx != index-1 && currentNode.next != nil {
		currentNode = currentNode.next
		idx++
	}

	if idx == index-1 {
		if currentNode.next != nil {
			currentNode.next = currentNode.next.next
		}
	}
}
