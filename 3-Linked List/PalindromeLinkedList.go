package main

func isPalindrome(head *ListNode) bool {
	mid := func(head *ListNode) *ListNode {
		slow, fast := head, head

		for fast != nil && fast.Next != nil {
			slow = slow.Next
			fast = fast.Next.Next
		}
		return slow
	}

	reverse := func(head *ListNode) *ListNode {
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

	startReverse := mid(head)
	leftList := head
	rightList := reverse(startReverse)
	for leftList != nil && rightList != nil {
		if leftList.Val != rightList.Val {
			return false
		}
		leftList = leftList.Next
		rightList = rightList.Next
	}
	return true
}
