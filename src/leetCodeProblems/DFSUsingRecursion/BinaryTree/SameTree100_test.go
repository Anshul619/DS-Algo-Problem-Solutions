package main

import "testing"

func TestIsSameTree(t *testing.T) {

	t.Run("test1", func(t *testing.T) {
		root := new(TreeNode) //Returns pointer to TreeNode object
		root.Val = 1
		root.Left = new(TreeNode)
		root.Left.Val = 2
		root.Right = new(TreeNode)
		root.Right.Val = 3

		root1 := new(TreeNode) //Returns pointer to TreeNode object
		root1.Val = 1
		root1.Left = new(TreeNode)
		root1.Left.Val = 2
		root1.Right = new(TreeNode)
		root1.Right.Val = 3

		if !isSameTree(root, root1) {
			t.Errorf("failed test1")
		}
	})

	t.Run("test2", func(t *testing.T) {
		root := new(TreeNode) //Returns pointer to TreeNode object
		root.Val = 1
		root.Left = new(TreeNode)
		root.Left.Val = 2

		root1 := new(TreeNode) //Returns pointer to TreeNode object
		root1.Val = 1
		root1.Right = new(TreeNode)
		root1.Right.Val = 2

		if isSameTree(root, root1) {
			t.Errorf("failed test2")
		}
	})
}
