package main

import "testing"

func TestSumNumbers(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		root := new(TreeNode) //Returns pointer to TreeNode object
		root.Val = 1
		root.Left = new(TreeNode)
		root.Left.Val = 2

		root.Right = new(TreeNode)
		root.Right.Val = 3

		if sumNumbers(root) != 25 {
			t.Fail()
		}
	})

	t.Run("test2", func(t *testing.T) {
		root := new(TreeNode) //Returns pointer to TreeNode object
		root.Val = 4
		root.Left = new(TreeNode)
		root.Left.Val = 9
		root.Left.Left = new(TreeNode)
		root.Left.Left.Val = 5
		root.Left.Right = new(TreeNode)
		root.Left.Right.Val = 1

		root.Right = new(TreeNode)
		root.Right.Val = 0

		if sumNumbers(root) != 1026 {
			t.Fail()
		}
	})
}
