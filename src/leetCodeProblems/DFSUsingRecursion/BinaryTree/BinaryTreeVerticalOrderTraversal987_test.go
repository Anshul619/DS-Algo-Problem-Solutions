package main

/*
- LeetCode - https://leetcode.com/problems/vertical-order-traversal-of-a-binary-tree/submissions/
*/

import (
	"reflect"
	"testing"
)

func TestVOrder1(t *testing.T) {
	root := new(TreeNode) //Returns pointer to TreeNode object
	root.Val = 1
	root.Left = new(TreeNode)
	root.Left.Val = 2
	root.Left.Left = new(TreeNode)
	root.Left.Left.Val = 4
	root.Left.Right = new(TreeNode)
	root.Left.Right.Val = 5

	root.Right = new(TreeNode)
	root.Right.Val = 3
	root.Right.Left = new(TreeNode)
	root.Right.Left.Val = 6
	root.Right.Right = new(TreeNode)
	root.Right.Right.Val = 7

	expectedOut := [][]int{{4}, {2}, {1, 5, 6}, {3}, {7}}

	if !reflect.DeepEqual(verticalTraversal(root), expectedOut) {
		t.Fatalf("TestVOrder Case Failed")
	}
}

func TestVOrder2(t *testing.T) {
	root := new(TreeNode) //Returns pointer to TreeNode object
	root.Val = 1
	root.Left = new(TreeNode)
	root.Left.Val = 2
	root.Left.Left = new(TreeNode)
	root.Left.Left.Val = 4
	root.Left.Right = new(TreeNode)
	root.Left.Right.Val = 6

	root.Right = new(TreeNode)
	root.Right.Val = 3
	root.Right.Left = new(TreeNode)
	root.Right.Left.Val = 5
	root.Right.Right = new(TreeNode)
	root.Right.Right.Val = 7

	expectedOut := [][]int{{4}, {2}, {1, 5, 6}, {3}, {7}}

	if !reflect.DeepEqual(verticalTraversal(root), expectedOut) {
		t.Fatalf("TestVOrder Case Failed")
	}
}

func TestVOrder3(t *testing.T) {
	root := new(TreeNode) //Returns pointer to TreeNode object
	root.Val = 0
	root.Left = new(TreeNode)
	root.Left.Val = 1
	root.Left.Right = new(TreeNode)
	root.Left.Right.Val = 2
	root.Left.Right.Left = new(TreeNode)
	root.Left.Right.Left.Val = 6
	root.Left.Right.Right = new(TreeNode)
	root.Left.Right.Right.Val = 3
	root.Left.Right.Right.Right = new(TreeNode)
	root.Left.Right.Right.Right.Val = 4
	root.Left.Right.Right.Right.Right = new(TreeNode)
	root.Left.Right.Right.Right.Right.Val = 5

	expectedOut := [][]int{{1, 6}, {0, 2}, {3}, {4}, {5}}

	if !reflect.DeepEqual(verticalTraversal(root), expectedOut) {
		t.Fatalf("TestVOrder Case Failed")
	}
}

func TestVOrder4(t *testing.T) {
	root := new(TreeNode) //Returns pointer to TreeNode object
	root.Val = 3
	root.Left = new(TreeNode)
	root.Left.Val = 1
	root.Left.Left = new(TreeNode)
	root.Left.Left.Val = 0
	root.Left.Right = new(TreeNode)
	root.Left.Right.Val = 2

	root.Right = new(TreeNode)
	root.Right.Val = 4
	root.Right.Left = new(TreeNode)
	root.Right.Left.Val = 2

	expectedOut := [][]int{{0}, {1}, {3, 2, 2}, {4}}

	if !reflect.DeepEqual(verticalTraversal(root), expectedOut) {
		t.Fatalf("TestVOrder Case Failed")
	}
}
