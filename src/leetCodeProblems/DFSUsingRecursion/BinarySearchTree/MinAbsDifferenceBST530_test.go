package main

/*
- Leetcode - https://leetcode.com/problems/minimum-absolute-difference-in-bst/description/
*/
import "testing"

func TestMinAbsDifference(t *testing.T) {
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
		root.Right.Val = 6

		if getMinimumDifference(root) != 1 {
			t.Fail()
		}
	})

	t.Run("test2", func(t *testing.T) {
		root := new(TreeNode) //Returns pointer to TreeNode object
		root.Val = 5
		root.Left = new(TreeNode)
		root.Left.Val = 4
		root.Right = new(TreeNode)
		root.Right.Val = 7

		if getMinimumDifference(root) != 1 {
			t.Fail()
		}
	})
}
