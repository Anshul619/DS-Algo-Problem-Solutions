package main

import "testing"

func TestDiameter(t *testing.T) {
	t.Run("test1", func(t *testing.T) {

		root := new(TreeNode) //Returns pointer to TreeNode object
		root.Val = 4
		root.Left = new(TreeNode)
		root.Left.Val = 2
		root.Left.Left = new(TreeNode)
		root.Left.Left.Val = 1
		root.Left.Right = new(TreeNode)
		root.Left.Right.Val = 3

		root.Right = new(TreeNode)
		root.Right.Val = 7
		root.Right.Left = new(TreeNode)
		root.Right.Left.Val = 6
		root.Right.Right = new(TreeNode)
		root.Right.Right.Val = 9

		if diameterOfBinaryTree(root) != 4 {
			t.Errorf("test failed")
		}

	})

	t.Run("test2", func(t *testing.T) {

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

		if diameterOfBinaryTree(root) != 3 {
			t.Errorf("test failed")
		}

	})
}
