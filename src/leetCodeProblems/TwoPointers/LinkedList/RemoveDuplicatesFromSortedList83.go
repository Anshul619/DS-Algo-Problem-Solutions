package main

/*
- Leetcode - https://leetcode.com/problems/remove-duplicates-from-sorted-list/description/
- Time - O(n)
- Space - O(1)
*/
// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	p := head

	for p.Next != nil {
		if p.Val == p.Next.Val {
			p.Next = p.Next.Next
		} else {
			p = p.Next
		}
	}
	return head
}
