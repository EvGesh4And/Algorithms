package main

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	fakeHead := &ListNode{}
	current := fakeHead
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			current.Next = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			list2 = list2.Next
		}
		current = current.Next
	}
	if list1 != nil {
		current.Next = list1
	}
	if list2 != nil {
		current.Next = list2
	}
	return fakeHead.Next
}
