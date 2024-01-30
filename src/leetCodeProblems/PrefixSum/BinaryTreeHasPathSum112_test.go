package main

import "testing"

func TestBinaryTreeHasPathSum112(t *testing.T) {

	t.Run("test1", func(t *testing.T) {
		root := new(TreeNode) //Returns pointer to TreeNode object
		root.Val = 5
		root.Left = new(TreeNode)
		root.Left.Val = 4
		root.Left.Left = new(TreeNode)
		root.Left.Left.Val = 11
		root.Left.Left.Left = new(TreeNode)
		root.Left.Left.Left.Val = 7
		root.Left.Left.Right = new(TreeNode)
		root.Left.Left.Right.Val = 2

		root.Right = new(TreeNode)
		root.Right.Val = 8
		root.Right.Left = new(TreeNode)
		root.Right.Left.Val = 13
		root.Right.Right = new(TreeNode)
		root.Right.Right.Val = 4
		root.Right.Right.Right = new(TreeNode)
		root.Right.Right.Right.Val = 1

		if !hasPathSum(root, 22) {
			t.Fail()
		}
	})

	t.Run("test2", func(t *testing.T) {
		root := new(TreeNode) //Returns pointer to TreeNode object
		root.Val = 1
		root.Left = new(TreeNode)
		root.Left.Val = 2

		root.Right = new(TreeNode)
		root.Right.Val = 3

		if hasPathSum(root, 5) {
			t.Fail()
		}
	})

	t.Run("test3", func(t *testing.T) {
		root := new(TreeNode) //Returns pointer to TreeNode object

		if !hasPathSum(root, 0) {
			t.Fail()
		}
	})
}
