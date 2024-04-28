package main

/*
- Leetcode - https://leetcode.com/problems/symmetric-tree/
- Time - O(n)
- Space - O(1)
*/

func isSymmetricUtil(node1, node2 *TreeNode) bool {
	if node1 == nil && node2 == nil {
		return true
	}

	if node1 == nil || node2 == nil || node1.Val != node2.Val {
		return false
	}

	return isSymmetricUtil(node1.Left, node2.Right) && isSymmetricUtil(node1.Right, node2.Left)
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return isSymmetricUtil(root.Left, root.Right)
}
