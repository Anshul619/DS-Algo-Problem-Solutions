package main

import (
	"math"
)

/**
 * LeetCode - https://leetcode.com/problems/count-complete-tree-nodes/
 *
 * Time Complexity - O(log2n )
 * Space Complexity - O(1)
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func powInt(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

func findDepth(root *TreeNode) int {
	depth := 0
	node := root

	for node != nil {
		depth++
		node = node.Left
	}
	return depth
}

func exists(index, depth int, root *TreeNode) bool {
	left := 1
	right := powInt(2, depth-1)

	node := root

	for i := 0; i < depth-1; i++ {
		mid := left + (right-left)/2

		if index <= mid {
			node = node.Left
			right = mid
		} else {
			node = node.Right
			left = mid + 1
		}
	}

	return node != nil
}

func countNodes(root *TreeNode) int {
	depth := findDepth(root)

	left := 1
	right := powInt(2, depth-1)

	for left <= right {
		mid := (left + right) / 2

		if exists(mid, depth, root) {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	out := 0

	for i := 0; i < depth-1; i++ {
		out += powInt(2, i)
	}

	out += left
	return out - 1
}
