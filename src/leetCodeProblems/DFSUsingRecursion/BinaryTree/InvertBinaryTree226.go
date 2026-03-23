package main

/*
- Leetcode - https://leetcode.com/problems/invert-binary-tree/description/
- Time - O(n) where n = number of nodes in the tree
- Space - O(1) + O(h) (recursion call stack) = O(h)
*/
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	// Swap in same instruction
	root.Left, root.Right = root.Right, root.Left

	invertTree(root.Left)
	invertTree(root.Right)

	return root
}
