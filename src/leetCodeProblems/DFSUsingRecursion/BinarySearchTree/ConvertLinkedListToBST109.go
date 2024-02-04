package main

/*
- LeetCode - https://leetcode.com/problems/convert-sorted-list-to-binary-search-tree/
- Time - O(n)
- Space - O(1)
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

var listNode *ListNode

func countNodes(head *ListNode) int {

	n := 0
	temp := head

	for temp != nil {
		n++
		temp = temp.Next
	}

	return n
}

func sortedListToBSTUtil(n int) *TreeNode {

	if n <= 0 {
		return nil
	}

	left := sortedListToBSTUtil(n / 2)

	treeNode := new(TreeNode)
	treeNode.Val = listNode.Val
	treeNode.Left = left

	listNode = listNode.Next

	treeNode.Right = sortedListToBSTUtil(n - n/2 - 1)

	return treeNode
}

func sortedListToBST(head *ListNode) *TreeNode {
	listNode = head
	return sortedListToBSTUtil(countNodes(head))
}

// func main() {
// 	head := new(ListNode)
// 	head.Val = -10
// 	head.Next = new(ListNode)
// 	head.Next.Val = -3
// 	head.Next.Next = new(ListNode)
// 	head.Next.Next.Val = 0
// 	head.Next.Next.Next = new(ListNode)
// 	head.Next.Next.Next.Val = 5
// 	head.Next.Next.Next.Next = new(ListNode)
// 	head.Next.Next.Next.Next.Val = 9

// 	printTree(sortedListToBST(head))
// }
