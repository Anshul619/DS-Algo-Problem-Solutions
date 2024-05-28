package main

import "testing"

func TestCountNodes(t *testing.T) {
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

	if countNodes(root) != 6 {
		t.Errorf("Test failed")
	}
}
