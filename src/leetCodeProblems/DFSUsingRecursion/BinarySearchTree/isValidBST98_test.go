package main

/*
- Leetcode - https://leetcode.com/problems/minimum-absolute-difference-in-bst/description/
*/
import "testing"

func TestIsValidBST(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		root := new(TreeNode) //Returns pointer to TreeNode object
		root.Val = 2
		root.Left = new(TreeNode)
		root.Left.Val = 1
		root.Right = new(TreeNode)
		root.Right.Val = 3

		expected := true

		if isValidBST(root) != expected {
			t.Fail()
		}
	})

	t.Run("test2", func(t *testing.T) {
		root := new(TreeNode) //Returns pointer to TreeNode object
		root.Val = 5
		root.Left = new(TreeNode)
		root.Left.Val = 1
		root.Right = new(TreeNode)
		root.Right.Val = 4
		root.Right.Left = new(TreeNode)
		root.Right.Left.Val = 3
		root.Right.Right = new(TreeNode)
		root.Right.Right.Val = 6

		expected := false

		if isValidBST(root) != expected {
			t.Fail()
		}
	})

	t.Run("test3", func(t *testing.T) {
		root := new(TreeNode) //Returns pointer to TreeNode object
		root.Val = 2
		root.Left = new(TreeNode)
		root.Left.Val = 2
		root.Right = new(TreeNode)
		root.Right.Val = 2

		expected := false

		if isValidBST(root) != expected {
			t.Fail()
		}
	})

	t.Run("test4", func(t *testing.T) {
		root := new(TreeNode) //Returns pointer to TreeNode object
		root.Val = 5
		root.Left = new(TreeNode)
		root.Left.Val = 4
		root.Right = new(TreeNode)
		root.Right.Val = 6
		root.Right.Left = new(TreeNode)
		root.Right.Left.Val = 3
		root.Right.Right = new(TreeNode)
		root.Right.Right.Val = 7

		expected := false

		if isValidBST(root) != expected {
			t.Fail()
		}
	})
}
