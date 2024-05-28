package main

/*
- LeetCode - https://leetcode.com/problems/binary-tree-level-order-traversal/description
*/

import (
	"reflect"
	"testing"
)

func TestLevelOrder(t *testing.T) {
	root := new(TreeNode) //Returns pointer to TreeNode object
	root.Val = 3
	root.Left = new(TreeNode)
	root.Left.Val = 9

	root.Right = new(TreeNode)
	root.Right.Val = 20
	root.Right.Left = new(TreeNode)
	root.Right.Left.Val = 15
	root.Right.Right = new(TreeNode)
	root.Right.Right.Val = 7

	expectedOut := [][]int{{3}, {9, 20}, {15, 7}}

	if !reflect.DeepEqual(levelOrder(root), expectedOut) {
		t.Fatalf("TestVOrder Case Failed")
	}
}
