package main

import (
	"reflect"
	"testing"
)

func TestRightSideView(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		root := new(TreeNode) //Returns pointer to TreeNode object
		root.Val = 1
		root.Left = new(TreeNode)
		root.Left.Val = 2
		root.Left.Right = new(TreeNode)
		root.Left.Right.Val = 5

		root.Right = new(TreeNode)
		root.Right.Val = 3
		root.Right.Right = new(TreeNode)
		root.Right.Right.Val = 4

		if !reflect.DeepEqual(rightSideView(root), []int{1, 3, 4}) {
			t.Fail()
		}
	})

	t.Run("test2", func(t *testing.T) {
		root := new(TreeNode) //Returns pointer to TreeNode object
		root.Val = 1
		root.Right = new(TreeNode)
		root.Right.Val = 3

		if !reflect.DeepEqual(rightSideView(root), []int{1, 3}) {
			t.Fail()
		}
	})

	t.Run("test3", func(t *testing.T) {
		if !reflect.DeepEqual(rightSideView(nil), []int{}) {
			t.Fail()
		}
	})
}
