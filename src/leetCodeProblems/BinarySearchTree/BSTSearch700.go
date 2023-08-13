package main

import "log"

/*
- LeetCode - https://leetcode.com/problems/search-in-a-binary-search-tree/
- Time - O(n)
- Space - O(1)
*/

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == val {
		return root
	}

	if root.Val < val {
		return searchBST(root.Right, val)
	}

	return searchBST(root.Left, val)
}

func main() {
	root := new(TreeNode)
	root.Val = 4
	root.Left = new(TreeNode)
	root.Left.Val = 2
	root.Right = new(TreeNode)
	root.Right.Val = 7
	root.Left.Left = new(TreeNode)
	root.Left.Left.Val = 1
	root.Left.Right = new(TreeNode)
	root.Left.Right.Val = 3

	//root = searchBST(root, 2)
	root = searchBST(root, 5)

	log.Println("root", root)
}
