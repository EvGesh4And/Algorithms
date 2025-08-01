package main

type node struct {
	val  int
	next *node
}

type MyLinkedList struct {
	len  int
	head *node
}

func Constructor() MyLinkedList {
	return MyLinkedList{0, &node{}}
}

func (this *MyLinkedList) Get(index int) int {
	if index >= 0 && index < this.len {
		curr := this.head
		for _ = range index + 1 {
			curr = curr.next
		}
		return curr.val
	}
	return -1
}

func (this *MyLinkedList) AddAtHead(val int) {
	this.AddAtIndex(0, val)
}

func (this *MyLinkedList) AddAtTail(val int) {
	this.AddAtIndex(this.len, val)
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index >= 0 && index <= this.len {
		curr := this.head
		for _ = range index {
			curr = curr.next
		}
		curr.next = &node{val, curr.next}
		this.len++
	}
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index >= 0 && index < this.len {
		curr := this.head
		for _ = range index {
			curr = curr.next
		}
		curr.next = curr.next.next
		this.len--
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
