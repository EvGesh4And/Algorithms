package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	fakeHead := &ListNode{Next: head}

	current := fakeHead

	for current.Next != nil && current.Next.Next != nil {
		first := current.Next
		second := current.Next.Next

		current.Next = second
		first.Next = second.Next
		second.Next = first
		current = first
	}
	return fakeHead.Next
}
