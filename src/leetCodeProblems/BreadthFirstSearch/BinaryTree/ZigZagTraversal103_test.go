package main

/*
- LeetCode - https://leetcode.com/problems/binary-tree-level-order-traversal/description
*/

import (
	"reflect"
	"testing"
)

func TestZigZagLevelOrder(t *testing.T) {
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

	expectedOut := [][]int{{3}, {20, 9}, {15, 7}}

	if !reflect.DeepEqual(zigzagLevelOrder(root), expectedOut) {
		t.Fatalf("TestZigZagLevelOrder Case Failed")
	}
}

func TestZigZagLevelOrder2(t *testing.T) {
	root := new(TreeNode) //Returns pointer to TreeNode object
	root.Val = 1
	root.Left = new(TreeNode)
	root.Left.Val = 2
	root.Left.Left = new(TreeNode)
	root.Left.Left.Val = 4

	root.Right = new(TreeNode)
	root.Right.Val = 3
	root.Right.Right = new(TreeNode)
	root.Right.Right.Val = 5

	expectedOut := [][]int{{1}, {3, 2}, {4, 5}}

	if !reflect.DeepEqual(zigzagLevelOrder(root), expectedOut) {
		t.Fatalf("TestZigZagLevelOrder2 Case Failed")
	}
}
