package main

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	past, current := head, head

	for current != nil {
		if current.Val != past.Val {
			past.Next = current
			past = current
		}
		current = current.Next
	}

	past.Next = nil // Обрываем возможную остаточную связь с дубликатом

	return head
}
