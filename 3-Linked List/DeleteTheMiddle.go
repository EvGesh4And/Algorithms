package main

// Более красивое решение как по мне
func deleteMiddle(head *ListNode) *ListNode {
	fakeHead := &ListNode{Next: head}
	slow, fast := fakeHead, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	slow.Next = slow.Next.Next
	return fakeHead.Next
}

// Имеет место, но не очень
func deleteMiddle(head *ListNode) *ListNode {

	if head.Next == nil {
		head = nil
		return head
	}

	fast, slow := head.Next.Next, head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	slow.Next = slow.Next.Next

	return head
}
