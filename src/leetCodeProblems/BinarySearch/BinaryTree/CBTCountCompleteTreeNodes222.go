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
	d := 0

	for root != nil && root.Left != nil {
		d++
		root = root.Left
	}
	return d
}

func exists(idx, d int, root *TreeNode) bool {
	l := 1
	r := powInt(2, d)

	for range d {
		m := l + (r-l)/2

		if idx <= m {
			root = root.Left
			r = m
		} else {
			root = root.Right
			l = m + 1
		}
	}
	return root != nil
}

func countNodes(root *TreeNode) int {
	d := findDepth(root)

	l := 1
	r := powInt(2, d)

	for l <= r {
		m := l + (r-l)/2

		if exists(m, d, root) {
			l = m + 1
		} else {
			r = m - 1
		}
	}

	out := 0

	for i := range d {
		out += powInt(2, i)
	}
	out += l
	return out - 1
}
