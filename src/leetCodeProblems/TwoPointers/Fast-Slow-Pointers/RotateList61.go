package main

/*
- LeetCode - https://leetcode.com/problems/rotate-list/
- Time - O(n)
- Space - O(1)
*/

func rotateRight(head *ListNode, k int) *ListNode {
	if k == 0 || head == nil {
		return head
	}
	total := 0
	next := head

	for next != nil {
		total++
		next = next.Next
	}

	k = k % total

	slow, fast := head, head

	for range k {
		fast = fast.Next
	}

	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}

	fast.Next = head
	newHead := slow.Next
	slow.Next = nil
	return newHead
}
