package main

func reverseList(head *ListNode) *ListNode {
    var left, carrent *ListNode
    carrent = head

    for carrent != nil {
        carrent.Next, carrent, left = left, carrent.Next, carrent
    }
    head = left

    return head
}
