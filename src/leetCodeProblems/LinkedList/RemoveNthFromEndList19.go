package main

/*
- Leetcode - https://leetcode.com/problems/remove-nth-node-from-end-of-list/description/
- Time - O(n)
- Space - O(1)
*/

func removeNthFromEnd(head *ListNode, n int) *ListNode {

	// This will be n+1 node from start.
	// Considering n+1 as we are initializing nodeToRemove & nodeToRemovePrev to head.
	// Once we fast reaches the end of list, nodeToRemove & nodeToRemovePrev would traverse by n-1
	// This is fine, as we are initializing nodeToRemove & nodeToRemovePrev to head
	fast := head

	for i := 0; i < n; i++ {
		fast = fast.Next
	}

	nodeToRemove, nodeToRemovePrev := head, head
	for fast != nil {
		nodeToRemovePrev = nodeToRemove
		nodeToRemove = nodeToRemove.Next
		fast = fast.Next
	}

	if nodeToRemove == head {
		return head.Next
	}

	nodeToRemovePrev.Next = nodeToRemove.Next
	return head
}
