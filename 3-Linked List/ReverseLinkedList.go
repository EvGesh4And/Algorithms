package main

func reverseList(head *ListNode) *ListNode {

	var past *ListNode
	current := head

	for current != nil {
		future := current.Next
		current.Next = past
		past = current
		current = future
	}
	return past
}

// Как вариант
func reverseList(head *ListNode) *ListNode {
	fakeHead := &ListNode{}
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = fakeHead.Next
		fakeHead.Next = curr
		curr = next
	}
	return fakeHead.Next
}
