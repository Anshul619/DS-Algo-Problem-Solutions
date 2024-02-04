package main

/*
- LeetCode - https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree/submissions/
*/

import (
	"testing"
)

func inOrderTraversal1(node *TreeNode, out *[]int) {
	if node == nil {
		return
	}
	inOrderTraversal1(node.Left, out)
	*out = append(*out, node.Val)
	inOrderTraversal1(node.Right, out)
}

func TestSortedArrayToBST(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		node := sortedArrayToBST([]int{-10, -3, 0, 5, 9})

		out := []int{}
		inOrderTraversal1(node, &out)
		if node.Val != 0 {
			t.Fail()
		}
	})
	t.Run("test2", func(t *testing.T) {
		node := sortedArrayToBST([]int{1, 3})

		if node.Val != 1 && node.Val != 3 {
			t.Fail()
		}
	})
}
