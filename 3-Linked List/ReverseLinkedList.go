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
