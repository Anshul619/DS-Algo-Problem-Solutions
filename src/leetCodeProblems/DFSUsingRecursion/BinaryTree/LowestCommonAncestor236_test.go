package main

/*
- Leetcode - https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/description/
*/
import (
	"testing"
)

func TestLowestCommonAncestor(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		root := new(TreeNode) //Returns pointer to TreeNode object
		root.Val = 3
		root.Left = new(TreeNode)
		root.Left.Val = 5
		root.Left.Left = new(TreeNode)
		root.Left.Left.Val = 6
		root.Left.Right = new(TreeNode)
		root.Left.Right.Val = 2
		root.Left.Right.Left = new(TreeNode)
		root.Left.Right.Left.Val = 7
		root.Left.Right.Right = new(TreeNode)
		root.Left.Right.Right.Val = 4

		root.Right = new(TreeNode)
		root.Right.Val = 1
		root.Right.Left = new(TreeNode)
		root.Right.Left.Val = 0
		root.Right.Right = new(TreeNode)
		root.Right.Right.Val = 8

		n := lowestCommonAncestor(root, &TreeNode{5, nil, nil}, &TreeNode{1, nil, nil})

		if n == nil || n.Val != 3 {
			t.Fail()
		}
	})

	t.Run("test2", func(t *testing.T) {
		root := new(TreeNode) //Returns pointer to TreeNode object
		root.Val = 3
		root.Left = new(TreeNode)
		root.Left.Val = 5
		root.Left.Left = new(TreeNode)
		root.Left.Left.Val = 6
		root.Left.Right = new(TreeNode)
		root.Left.Right.Val = 2
		root.Left.Right.Left = new(TreeNode)
		root.Left.Right.Left.Val = 7
		root.Left.Right.Right = new(TreeNode)
		root.Left.Right.Right.Val = 4

		root.Right = new(TreeNode)
		root.Right.Val = 1
		root.Right.Left = new(TreeNode)
		root.Right.Left.Val = 0
		root.Right.Right = new(TreeNode)
		root.Right.Right.Val = 8

		n := lowestCommonAncestor(root, &TreeNode{5, nil, nil}, &TreeNode{4, nil, nil})

		if n == nil || n.Val != 5 {
			t.Fail()
		}
	})

	t.Run("test3", func(t *testing.T) {
		root := new(TreeNode) //Returns pointer to TreeNode object
		root.Val = 1
		root.Left = new(TreeNode)
		root.Left.Val = 2

		n := lowestCommonAncestor(root, &TreeNode{1, nil, nil}, &TreeNode{2, nil, nil})

		if n == nil || n.Val != 1 {
			t.Fail()
		}
	})
}
