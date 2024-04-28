package main

/**
- LeetCode - https://leetcode.com/problems/minimum-depth-of-binary-tree/
- Time - O(n)
- Space - O(1)
*/

func calculateHeight(node *TreeNode, height int, maxHeight *int) {
	if node == nil {
		return
	}
	if height > *maxHeight {
		*maxHeight = height
	}

	calculateHeight(node.Left, height+1, maxHeight)
	calculateHeight(node.Right, height+1, maxHeight)
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	out := 0
	calculateHeight(root, 0, &out)
	return out + 1
}
