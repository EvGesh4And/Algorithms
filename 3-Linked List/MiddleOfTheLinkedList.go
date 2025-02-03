package main

func middleNode(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	len := 1
	currentNode := head

	for currentNode.Next != nil {
		len++
		currentNode = currentNode.Next
	}

	currentNode = head

	for i := 0; i < len/2; i++ {
		currentNode = currentNode.Next
	}
	return currentNode
}
