package main

/*
- LeetCode - https://leetcode.com/problems/rotate-list/
- Time - O(n)
- Space - O(1)
*/

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	if k == 0 {
		return head
	}

	total := 0
	temp := head

	// count list
	for temp != nil {
		total++
		temp = temp.Next
	}

	if k > total {
		k = k % total
	}

	slow, fast := head, head

	// Move fast by leftK
	for i := 0; i < k; i++ {
		if fast.Next != nil {
			fast = fast.Next
		} else {
			return head
		}

	}

	// Move fast and slow till end of list
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}

	fast.Next = head
	newHead := slow.Next
	slow.Next = nil
	return newHead
}
