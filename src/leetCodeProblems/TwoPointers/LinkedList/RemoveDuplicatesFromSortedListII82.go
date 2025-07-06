package main

/*
- Leetcode - https://leetcode.com/problems/remove-duplicates-from-sorted-list-ii/description
- Time - O(n)
- Space - O(1)
*/
func deleteDuplicatesII(head *ListNode) *ListNode {
	// This dummy head is helpful, to handle null head
	dummy := &ListNode{Next: head} // Dummy node to simplify head deletion
	prev := dummy
	cur := head

	for cur != nil && cur.Next != nil {

		// Detect the start of duplicates
		if cur.Val == cur.Next.Val {

			// Move to the last node in this duplicate group
			for cur.Next != nil && cur.Val == cur.Next.Val {
				cur = cur.Next
			}

			// Skip the entire duplicate group
			prev.Next = cur.Next
		} else {

			// No duplication, move prev forward
			prev = cur
		}

		cur = cur.Next

	}
	return dummy.Next
}
