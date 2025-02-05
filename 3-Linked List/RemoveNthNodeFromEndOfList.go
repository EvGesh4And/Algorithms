package main

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// Фэйковый узел
	head = &ListNode{0, head}

	p1, p2 := head, head

	for i := 0; i < n; i++ {
		p2 = p2.Next
	}

	for p2.Next != nil {
		p1 = p1.Next
		p2 = p2.Next
	}

	p1.Next = p1.Next.Next

	return head.Next
}
