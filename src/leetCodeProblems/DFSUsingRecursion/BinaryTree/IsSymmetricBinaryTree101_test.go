package main

import "testing"

func TestIsSymmetric(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		root := new(TreeNode) //Returns pointer to TreeNode object
		root.Val = 1
		root.Left = new(TreeNode)
		root.Left.Val = 2
		root.Left.Left = new(TreeNode)
		root.Left.Left.Val = 3
		root.Left.Right = new(TreeNode)
		root.Left.Right.Val = 4

		root.Right = new(TreeNode)
		root.Right.Val = 2

		root.Right.Left = new(TreeNode)
		root.Right.Left.Val = 4
		root.Right.Right = new(TreeNode)
		root.Right.Right.Val = 3

		if !isSymmetric(root) {
			t.Errorf("test failed")
		}

	})
}
