package main

type node struct {
	val  int
	next *node
}

type MyLinkedList struct {
	head *node
}

func Constructor() MyLinkedList {
	return MyLinkedList{&node{}}
}

func (this *MyLinkedList) Get(index int) int {
	curr := this.head
	for pos := -1; pos != index; pos++ {
		curr = curr.next
		if curr == nil {
			return -1
		}
	}
	return curr.val
}

func (this *MyLinkedList) AddAtHead(val int) {
	this.head.next = &node{val, this.head.next}
}

func (this *MyLinkedList) AddAtTail(val int) {
	curr := this.head
	for curr.next != nil {
		curr = curr.next
	}
	curr.next = &node{val, nil}
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	curr := this.head
	for pos := -1; pos != index-1; pos++ {
		curr = curr.next
		if curr == nil {
			return
		}
	}
	curr.next = &node{val, curr.next}
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	curr := this.head
	for pos := -1; pos != index-1; pos++ {
		curr = curr.next
		if curr == nil {
			return
		}
	}
	if curr.next != nil {
		curr.next = curr.next.next
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
