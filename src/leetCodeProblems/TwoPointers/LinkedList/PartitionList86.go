package main

/*
- Leetcode - https://leetcode.com/problems/partition-list
- Time - O(n)
- Space - O(1)
*/
func partition(head *ListNode, x int) *ListNode {
	sDummy, lDummy := &ListNode{}, &ListNode{}
	small, large := sDummy, lDummy

	for head != nil {
		if head.Val < x {
			small.Next = head
			small = small.Next
		} else {
			large.Next = head
			large = large.Next
		}

		head = head.Next
	}

	large.Next = nil
	small.Next = lDummy.Next

	return sDummy.Next
}
