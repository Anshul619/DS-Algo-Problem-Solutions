package main

/*
- LeetCode - https://leetcode.com/problems/reverse-linked-list/
- TimeComplexity - O(N)
- SpaceComplexity - O(1)
*/

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	cur := head

	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}
