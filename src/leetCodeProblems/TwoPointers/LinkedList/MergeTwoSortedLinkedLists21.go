package main

/*
- LeetCode - https://leetcode.com/problems/merge-two-sorted-lists/description/
- Time - O(m+n)
- Space - O(1)
*/
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	prev := dummyHead

	for list1 != nil || list2 != nil {
		if list1 != nil && list2 != nil {
			if list1.Val > list2.Val {
				prev.Next = list2
				list2 = list2.Next
			} else {
				prev.Next = list1
				list1 = list1.Next
			}

			prev = prev.Next
		} else if list1 != nil {
			prev.Next = list1
			list1 = nil
		} else {
			prev.Next = list2
			list2 = nil
		}
	}
	return dummyHead.Next
}
