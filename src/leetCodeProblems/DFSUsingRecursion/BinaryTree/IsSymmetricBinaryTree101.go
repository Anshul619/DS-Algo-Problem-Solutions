package main

/*
- Leetcode - https://leetcode.com/problems/symmetric-tree/
- Time - O(n)
- Space - O(1)
*/

func dfs(node1, node2 *TreeNode) bool {
	if node1 == nil && node2 == nil {
		return true
	}
	if node1 == nil {
		return false
	}
	if node2 == nil {
		return false
	}

	if node1.Val != node2.Val {
		return false
	}

	return dfs(node1.Left, node2.Right) && dfs(node1.Right, node2.Left)
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return dfs(root.Left, root.Right)
}
