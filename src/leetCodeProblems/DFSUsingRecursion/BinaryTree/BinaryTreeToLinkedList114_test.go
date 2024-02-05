package main

/*
- Leetcode - https://leetcode.com/problems/flatten-binary-tree-to-linked-list/description/
- Time - O(n)
- Space - O(1)
*/
import (
	"testing"
)

func TestFlattenBT(t *testing.T) {
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
		root.Right.Val = 5
		root.Right.Right = new(TreeNode)
		root.Right.Right.Val = 6

		flatten(root)

		printInOrder(root)
	})

	t.Run("test2", func(t *testing.T) {
		root := new(TreeNode) //Returns pointer to TreeNode object
		root.Val = 1
		root.Left = new(TreeNode)
		root.Left.Val = 2
		root.Left.Left = new(TreeNode)
		root.Left.Left.Val = 3
		root.Left.Right = new(TreeNode)
		root.Left.Right.Val = 4
		root.Left.Left.Left = new(TreeNode)
		root.Left.Left.Left.Val = 5

		flatten(root)
	})
}
