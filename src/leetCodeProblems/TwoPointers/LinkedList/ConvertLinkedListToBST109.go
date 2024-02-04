package main

/*
- LeetCode - https://leetcode.com/problems/convert-sorted-list-to-binary-search-tree/
- Time - O(nlogn)
- Space - O(1)
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return &TreeNode{head.Val, nil, nil}
	}
	prevSlow, slow, fast := head, head, head
	for fast != nil && fast.Next != nil {
		prevSlow = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	node := new(TreeNode)
	node.Val = slow.Val
	prevSlow.Next = nil
	node.Left = sortedListToBST(head)
	node.Right = sortedListToBST(slow.Next)
	return node
}
