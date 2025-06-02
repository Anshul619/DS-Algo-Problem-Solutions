package main

import (
	"reflect"
	"testing"
)

func TestPathSum113(t *testing.T) {
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
	root.Right.Right.Left = new(TreeNode)
	root.Right.Right.Left.Val = 5
	root.Right.Right.Right = new(TreeNode)
	root.Right.Right.Right.Val = 1

	if !reflect.DeepEqual(pathSum1(root, 22), [][]int{{5, 4, 11, 2}, {5, 8, 4, 5}}) {
		t.Error()
	}
}
