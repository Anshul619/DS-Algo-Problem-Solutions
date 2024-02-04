package main

/*
- LeetCode - https://leetcode.com/problems/swapping-nodes-in-a-linked-list/description/
- Time - O(n)
- Space - O(1)
*/
func swapNodes(head *ListNode, k int) *ListNode {
	kNodeStart, fast := head, head

	// fast will be n+1 node from start.
	// We would initialize kNodeEnd to head.
	// Once fast reaches the end of list, kNodeEnd would traverse by n-1
	for i := 0; i < k; i++ {
		kNodeStart = fast
		fast = fast.Next
	}

	kNodeEnd := head
	for fast != nil {
		kNodeEnd = kNodeEnd.Next
		fast = fast.Next
	}

	// swap elements
	temp := kNodeStart.Val
	kNodeStart.Val = kNodeEnd.Val
	kNodeEnd.Val = temp
	return head
}
