package main

/*
- Leetcode - https://leetcode.com/problems/minimum-absolute-difference-in-bst/description/
- Similar Question - https://leetcode.com/problems/minimum-distance-between-bst-nodes/description/
- Time - O(n)
- Space - O(h)
*/
import (
	"math"
)

func checkMinDiff(node *TreeNode, lastVal, minDiff *int) {
	if node == nil {
		return
	}

	checkMinDiff(node.Left, lastVal, minDiff)

	if *lastVal != -1 && node.Val-*lastVal < *minDiff {
		*minDiff = node.Val - *lastVal

	}

	*lastVal = node.Val

	checkMinDiff(node.Right, lastVal, minDiff)
}

func getMinimumDifference(root *TreeNode) int {
	out, lastVal := math.MaxInt, -1

	checkMinDiff(root, &lastVal, &out)
	return out
}
