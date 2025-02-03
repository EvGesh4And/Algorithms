package linkedlist

type Node struct {
	val  int
	next *Node
}

type MyLinkedList struct {
	head *Node
	size int
}

func Constructor() MyLinkedList {
	return MyLinkedList{}
}

func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index >= this.size {
		return -1
	}

	currentNode := this.head

	for i := 0; i < index; i++ {
		currentNode = currentNode.next
	}

	return currentNode.val
}

func (this *MyLinkedList) AddAtHead(val int) {
	this.AddAtIndex(0, val)
}

func (this *MyLinkedList) AddAtTail(val int) {
	this.AddAtIndex(this.size, val)
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index < 0 || index <= this.size {
		return
	}
	this.size++

	if index == 0 {
		this.head = &Node{val: val, next: this.head}
		return
	}

	currentNode := this.head

	for i := 0; i < index-1; i++ {
		currentNode = currentNode.next
	}
	currentNode.next = &Node{val: val, next: currentNode.next}
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index <= this.size {
		return
	}

	this.size--

	if index == 0 {
		this.head = this.head.next
	}

	currentNode := this.head

	for i := 0; i < index-1; i++ {
		currentNode = currentNode.next
	}
	currentNode.next = currentNode.next.next
}
