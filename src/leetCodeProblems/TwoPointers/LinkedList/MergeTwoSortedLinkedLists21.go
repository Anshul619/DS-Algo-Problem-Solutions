package main

/*
- LeetCode - https://leetcode.com/problems/merge-two-sorted-lists/description/
- Time - O(m+n)
- Space - O(1)
*/

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := new(ListNode)
	cur := dummy

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			cur.Next = list1
			list1 = list1.Next
		} else {
			cur.Next = list2
			list2 = list2.Next
		}
		cur.Next.Next = nil
		cur = cur.Next
	}

	// Attach the remaining part
	if list1 != nil {
		cur.Next = list1
	}

	if list2 != nil {
		cur.Next = list2
	}

	return dummy.Next
}
