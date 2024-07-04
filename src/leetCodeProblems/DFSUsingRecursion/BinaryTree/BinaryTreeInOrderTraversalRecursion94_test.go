package main

import (
	"reflect"
	"testing"
)

func TestSample(t *testing.T) {

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

		if !reflect.DeepEqual(inorderTraversal(root), []int{3, 7, 9, 15, 20}) {
			t.Errorf("failed")
		}
	})
}
