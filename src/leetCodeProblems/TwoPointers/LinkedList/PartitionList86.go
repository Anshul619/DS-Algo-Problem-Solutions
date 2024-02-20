package main

/*
- Leetcode - https://leetcode.com/problems/partition-list
- Time - O(n)
- Space - O(1)
*/
func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return head
	}

	dummyHead := new(ListNode)
	lower := dummyHead

	prev := dummyHead
	prev.Next = head

	cur := prev.Next

	for cur != nil {
		if cur.Val < x {

			// Special case when first element of list, is slower than x
			if lower.Next == cur {
				lower = cur
				prev = cur
				cur = cur.Next
				continue
			}

			tempLowerNext := lower.Next

			// Move cur as next to lower and adjust prev.Next
			lower.Next = cur
			prev.Next = cur.Next
			cur.Next = tempLowerNext

			lower = lower.Next
			cur = prev.Next
		} else {
			prev = cur
			cur = cur.Next
		}
	}
	return dummyHead.Next
}
