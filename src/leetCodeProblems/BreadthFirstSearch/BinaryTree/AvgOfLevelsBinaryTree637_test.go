package main

/*
- Leetcode - https://leetcode.com/problems/average-of-levels-in-binary-tree
*/
import (
	"reflect"
	"testing"
)

func TestAverageOfLevels(t *testing.T) {

	t.Run("test1", func(t *testing.T) {
		root := new(TreeNode) //Returns pointer to TreeNode object
		root.Val = 3
		root.Left = new(TreeNode)
		root.Left.Val = 9

		root.Right = new(TreeNode)
		root.Right.Val = 20
		root.Right.Left = new(TreeNode)
		root.Right.Left.Val = 15
		root.Right.Right = new(TreeNode)
		root.Right.Right.Val = 7

		if !reflect.DeepEqual(averageOfLevels(root), []float64{3.00000, 14.50000, 11.00000}) {
			t.Fail()
		}
	})

	t.Run("test2", func(t *testing.T) {
		root := new(TreeNode) //Returns pointer to TreeNode object
		root.Val = 3
		root.Left = new(TreeNode)
		root.Left.Val = 9
		root.Left.Left = new(TreeNode)
		root.Left.Left.Val = 15
		root.Left.Right = new(TreeNode)
		root.Left.Right.Val = 7

		root.Right = new(TreeNode)
		root.Right.Val = 20

		if !reflect.DeepEqual(averageOfLevels(root), []float64{3.00000, 14.50000, 11.00000}) {
			t.Fail()
		}
	})
}
