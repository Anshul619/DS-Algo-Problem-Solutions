package main

import "testing"

func TestMaxDepth(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
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

		if maxDepth(root) != 3 {
			t.Errorf("test failed")
		}

	})
	t.Run("test2", func(t *testing.T) {
		root := new(TreeNode) //Returns pointer to TreeNode object
		root.Val = 1
		root.Right = new(TreeNode)
		root.Right.Val = 2

		if maxDepth(root) != 2 {
			t.Errorf("test failed")
		}

	})
}
