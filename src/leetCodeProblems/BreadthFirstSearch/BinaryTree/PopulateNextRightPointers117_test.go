package main

import "testing"

func TestConnect1(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		root := new(Node) //Returns pointer to TreeNode object
		root.Val = 1
		root.Left = new(Node)
		root.Left.Val = 2
		root.Left.Left = new(Node)
		root.Left.Left.Val = 4
		root.Left.Right = new(Node)
		root.Left.Right.Val = 5

		root.Right = new(Node)
		root.Right.Val = 3
		root.Right.Left = new(Node)
		root.Right.Left.Val = 6
		root.Right.Right = new(Node)
		root.Right.Right.Val = 7

		connect1(root)
	})
}
