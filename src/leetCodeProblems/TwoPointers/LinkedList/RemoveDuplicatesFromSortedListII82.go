package main

/*
- Leetcode - https://leetcode.com/problems/remove-duplicates-from-sorted-list-ii/description
- Time - O(n)
- Space - O(1)
*/
func deleteDuplicatesII(head *ListNode) *ListNode {
	// Dummy node simplifies removal at head
	dummy := &ListNode{}
	dummy.Next = head

	prev := dummy

	cur := head

	// Iterate through the list
	for cur != nil {

		// Check if current has duplicates by comparing with next node
		if cur.Next != nil && cur.Val == cur.Next.Val {

			// Skip all nodes with the same value
			for cur.Next != nil && cur.Val == cur.Next.Val {
				cur = cur.Next
			}

			// Link prev to the first node after duplicates. No need to move prev
			prev.Next = cur.Next
		} else {
			// Move prev to its next node
			prev = prev.Next
		}

		// Move cur to its next current
		cur = cur.Next
	}
	return dummy.Next
}
