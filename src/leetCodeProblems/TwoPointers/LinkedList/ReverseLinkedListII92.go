package main

/*
- Leetcode - https://leetcode.com/problems/reverse-linked-list-ii/description/
- Time - O(n)
- Space - O(1)
*/

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left == right {
		return head
	}

	if head == nil {
		return head
	}

	// This dummy head is helpful, to handle null head
	dummyHead := new(ListNode)
	prev := dummyHead

	prev.Next = head

	// Go to left pointer
	for i := 0; i < left-1; i++ {
		prev = prev.Next
	}

	cur := prev.Next

	for i := 0; i < right-left; i++ {
		t := prev.Next
		prev.Next = cur.Next
		cur.Next = cur.Next.Next
		prev.Next.Next = t
	}

	return dummyHead.Next
}
