package main

import "leetcode-solutions/pkg/common"

type ListNode = common.ListNode

// Medium: DeleteTheMiddleNodeOfALinkedList
// Solution for delete-the-middle-node-of-a-linked-list (medium)
func deleteMiddle(head *ListNode) *ListNode {
	posMap := make(map[int]*ListNode)
	current := head
	idx := 0

	for current != nil {
		posMap[idx] = current
		current = current.Next

		idx++
	}

	// last idx += 1 before exist the loop, so idx here is length `n`
	middlePos := idx / 2

	if middlePos != 0 {
		prevNode := posMap[middlePos-1]
		nextNode := posMap[middlePos+1]

		prevNode.Next = nextNode
		return head
	} else {
		return nil
	}
}

func deleteMiddle2(head *ListNode) *ListNode {
	if head.Next == nil {
		return nil
	}

	slow := head
	fast := head.Next

	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	slow.Next = slow.Next.Next
	return head
}
