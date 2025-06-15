package main

/*
- Leetcode - https://leetcode.com/problems/diameter-of-binary-tree/
- Time - O(n)
- Space - O(h)
*/

func diameterOfBinaryTreeDFS(node *TreeNode, out *int) int {
	if node == nil {
		return 0
	}

	left := diameterOfBinaryTreeDFS(node.Left, out)
	right := diameterOfBinaryTreeDFS(node.Right, out)

	if *out < (left + right) {
		*out = left + right
	}

	if left > right {
		return left + 1
	}

	return right + 1
}

func diameterOfBinaryTree(root *TreeNode) int {
	out := 0
	diameterOfBinaryTreeDFS(root, &out)
	return out
}
