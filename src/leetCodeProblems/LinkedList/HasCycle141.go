/*
- LeetCode - https://leetcode.com/problems/linked-list-cycle/
*/
package main

func hasCycle(head *ListNode) bool {

	nodesMap := make(map[*ListNode]bool)

	if head == nil {
		return false
	}

	nodesMap[head] = true

	for head.Next != nil {

		head = head.Next

		if _, ok := nodesMap[head]; ok {
			return true
		}

		nodesMap[head] = true
	}

	return false
}
