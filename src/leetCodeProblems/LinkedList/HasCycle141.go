/*
- LeetCode - https://leetcode.com/problems/linked-list-cycle/
*/
package main

func hasCycle(head *ListNode) bool {

	if head == nil {
		return false
	}

	slow := head
	fast := head

	for slow != nil && fast != nil && fast.Next != nil {

		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			return true
		}
	}

	return false
}
