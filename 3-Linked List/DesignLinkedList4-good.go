package main

type Node struct {
	val  int
	next *Node
}

type MyLinkedList struct {
	head *Node
}

func Constructor() MyLinkedList {
	return MyLinkedList{head: &Node{}}
}

func (this *MyLinkedList) Get(index int) int {
	pos := -1
	currNode := this.head
	for pos != index {
		currNode = currNode.next
		pos++
		if currNode == nil {
			return -1
		}
	}
	return currNode.val
}

func (this *MyLinkedList) AddAtHead(val int) {
	this.head.next = &Node{val: val, next: this.head.next}
}

func (this *MyLinkedList) AddAtTail(val int) {
	currNode := this.head
	for currNode.next != nil {
		currNode = currNode.next
	}
	currNode.next = &Node{val: val, next: nil}
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	pos := -1
	currNode := this.head
	for pos != index-1 {
		currNode = currNode.next
		pos++
		if currNode == nil {
			return
		}
	}
	currNode.next = &Node{val: val, next: currNode.next}
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	pos := -1
	currNode := this.head
	for pos != index-1 {
		currNode = currNode.next
		pos++
		if currNode == nil {
			return
		}
	}
	if currNode.next != nil {
		currNode.next = currNode.next.next
	}
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
