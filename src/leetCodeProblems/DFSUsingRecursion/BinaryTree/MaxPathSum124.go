package main

/*
- Leetcode - https://leetcode.com/problems/binary-tree-maximum-path-sum/description/?envType=study-plan-v2&envId=top-interview-150
- time - O(N)
- space - O(H)
*/
import "math"

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxPathSumDFS(node *TreeNode, max *int) int {
	if node == nil {
		return 0
	}

	leftGain := maxInt(0, maxPathSumDFS(node.Left, max))
	rightGain := maxInt(0, maxPathSumDFS(node.Right, max))

	*max = maxInt(*max, leftGain+node.Val+rightGain)

	return node.Val + maxInt(leftGain, rightGain)
}

func maxPathSum(root *TreeNode) int {
	max := math.MinInt
	maxPathSumDFS(root, &max)
	return max
}
