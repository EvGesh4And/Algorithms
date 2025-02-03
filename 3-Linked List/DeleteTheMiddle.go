package main

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
