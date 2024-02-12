package main

/*
- LeetCode - https://leetcode.com/problems/reverse-linked-list/
- TimeComplexity - O(n)
- SpaceComplexity - O(n)
*/

// func reverseBetweenUtil(node *ListNode, left int, right int, reverse bool) *ListNode {
// 	if node == nil || node.Next == nil {
// 		return node
// 	}

// 	if !reverse {
// 		return node
// 	}

// 	var new *ListNode

// 	if node.Val == left {
// 		new = reverseBetweenUtil(node.Next, left, right, true)
// 	}

// 	if node.Val == right {
// 		return node
// 	}

// }

// func reverseBetween(head *ListNode, left int, right int) *ListNode {
// 	return reverseBetweenUtil(head, left, right, false)
// }
