package main

import (
	"testing"
)

func TestBSTIterator(t *testing.T) {

	t.Run("test1", func(t *testing.T) {
		root := new(TreeNode) //Returns pointer to TreeNode object
		root.Val = 7
		root.Left = new(TreeNode)
		root.Left.Val = 3

		root.Right = new(TreeNode)
		root.Right.Val = 15
		root.Right.Left = new(TreeNode)
		root.Right.Left.Val = 9
		root.Right.Right = new(TreeNode)
		root.Right.Right.Val = 20

		bst := Constructor(root)

		if bst.Next() != 3 {
			t.Fail()
		}

		if bst.Next() != 7 {
			t.Fail()
		}

		if !bst.HasNext() {
			t.Fail()
		}

		if bst.Next() != 9 {
			t.Fail()
		}

		if !bst.HasNext() {
			t.Fail()
		}

		if bst.Next() != 15 {
			t.Fail()
		}

		if !bst.HasNext() {
			t.Fail()
		}

		if bst.Next() != 20 {
			t.Fail()
		}

		if bst.HasNext() {
			t.Fail()
		}

	})
}
