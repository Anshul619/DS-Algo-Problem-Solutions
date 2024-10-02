package main

/*
- LeetCode - https://leetcode.com/problems/merge-two-sorted-lists/description/
- Time - O(m+n)
- Space - O(1)
*/

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var outHead, outNext *ListNode

	for list1 != nil || list2 != nil {
		var nodeToInsert *ListNode

		switch {
		case list1 == nil:
			nodeToInsert = list2
			list2 = nil
		case list2 == nil:
			nodeToInsert = list1
			list1 = nil
		default:
			if list1.Val < list2.Val {
				nodeToInsert = list1
				list1 = list1.Next
			} else {
				nodeToInsert = list2
				list2 = list2.Next
			}
		}

		if outHead != nil {
			outNext.Next = nodeToInsert
		} else {
			outHead = nodeToInsert
		}

		outNext = nodeToInsert
	}

	return outHead
}
