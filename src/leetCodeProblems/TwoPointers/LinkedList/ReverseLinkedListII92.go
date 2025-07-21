package main

/*
- Leetcode - https://leetcode.com/problems/reverse-linked-list-ii/description/
- Time - O(n)
- Space - O(1)
*/

// Note - Both left, right are 1-indexed
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil || left == right {
		return head
	}

	// dummy prev node
	dummy := &ListNode{}
	dummy.Next = head
	prev := dummy

	// Move `prev` to the node before reversal starts
	for i := 1; i < left; i++ {
		prev = prev.Next
	}

	start := prev.Next // start of the reversal list
	move := start.Next // node to be moved

	// move then (for right-left times)
	for i := 0; i < right-left; i++ {
		start.Next = move.Next
		move.Next = prev.Next
		prev.Next = move
		move = start.Next
	}

	return dummy.Next
}
