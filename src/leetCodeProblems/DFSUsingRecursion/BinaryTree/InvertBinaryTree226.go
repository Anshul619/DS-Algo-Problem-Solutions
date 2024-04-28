package main

/*
- Leetcode - https://leetcode.com/problems/invert-binary-tree/description/
- Time - O(n)
- Space - O(1)
*/

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	node := new(TreeNode)
	node.Val = root.Val
	node.Left = invertTree(root.Right)
	node.Right = invertTree(root.Left)
	return node
}
