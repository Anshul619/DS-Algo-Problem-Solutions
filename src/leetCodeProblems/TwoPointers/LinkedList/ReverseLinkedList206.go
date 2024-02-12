package main

/*
- LeetCode - https://leetcode.com/problems/reverse-linked-list/
- TimeComplexity - O(N)
- SpaceComplexity - O(1)
*/

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	next := head.Next
	prev := head

	for next != nil {

		temp := next.Next
		next.Next = head
		prev.Next = temp
		head = next
		next = temp
	}
	return head
}
