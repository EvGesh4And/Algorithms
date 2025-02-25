package main

type Node struct {
	val  int
	next *Node
}

type MyLinkedList struct {
	head    *Node
	lenList int
}

func Constructor() MyLinkedList {
	return MyLinkedList{head: &Node{}, lenList: 0}
}

func (this *MyLinkedList) Get(index int) int {
	if index < this.lenList {
		currNode := this.head
		for i := 0; i <= index; i++ {
			currNode = currNode.next
		}
		return currNode.val
	} else {
		return -1
	}
}

func (this *MyLinkedList) AddAtHead(val int) {
	this.head.next = &Node{val: val, next: this.head.next}
	this.lenList++
}

func (this *MyLinkedList) AddAtTail(val int) {
	currNode := this.head
	for currNode.next != nil {
		currNode = currNode.next
	}
	currNode.next = &Node{val: val, next: nil}
	this.lenList++
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index <= this.lenList {
		currNode := this.head
		for i := 0; i < index; i++ {
			currNode = currNode.next
		}
		currNode.next = &Node{val: val, next: currNode.next}
		this.lenList++
	}
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < this.lenList {
		currNode := this.head
		for i := 0; i < index; i++ {
			currNode = currNode.next
		}

		currNode.next = currNode.next.next
		this.lenList--
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
