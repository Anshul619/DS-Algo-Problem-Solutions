package main

import (
	"reflect"
	"testing"
)

func Test_Infection1(t *testing.T) {
	root := new(TreeNode) //Returns pointer to TreeNode object
	root.Val = 1
	root.Left = new(TreeNode)
	root.Left.Val = 5
	root.Left.Right = new(TreeNode)
	root.Left.Right.Val = 4
	root.Left.Right.Left = new(TreeNode)
	root.Left.Right.Left.Val = 9
	root.Left.Right.Right = new(TreeNode)
	root.Left.Right.Right.Val = 2

	root.Right = new(TreeNode)
	root.Right.Val = 3
	root.Right.Left = new(TreeNode)
	root.Right.Left.Val = 10
	root.Right.Right = new(TreeNode)
	root.Right.Right.Val = 6

	if !reflect.DeepEqual(amountOfTime(root, 3), 4) {
		t.Fatalf("Test_Infection1 Case Failed")
	}
}

func Test_Infection2(t *testing.T) {
	root := new(TreeNode) //Returns pointer to TreeNode object
	root.Val = 1
	root.Left = new(TreeNode)
	root.Left.Val = 2
	root.Left.Left = new(TreeNode)
	root.Left.Left.Val = 3
	root.Left.Left.Left = new(TreeNode)
	root.Left.Left.Left.Val = 4
	root.Left.Left.Left.Left = new(TreeNode)
	root.Left.Left.Left.Left.Val = 5

	if !reflect.DeepEqual(amountOfTime(root, 3), 2) {
		t.Fatalf("Test_Infection2 Case Failed")
	}
}

func Test_Infection3(t *testing.T) {
	root := new(TreeNode) //Returns pointer to TreeNode object
	root.Val = 1
	root.Left = new(TreeNode)
	root.Left.Val = 2
	root.Left.Left = new(TreeNode)
	root.Left.Left.Val = 3
	root.Left.Left.Left = new(TreeNode)
	root.Left.Left.Left.Val = 4
	root.Left.Left.Left.Left = new(TreeNode)
	root.Left.Left.Left.Left.Val = 5

	if !reflect.DeepEqual(amountOfTime(root, 2), 3) {
		t.Fatalf("Test_Infection3 Case Failed")
	}
}

func Test_Infection4(t *testing.T) {
	root := new(TreeNode) //Returns pointer to TreeNode object
	root.Val = 1
	root.Left = new(TreeNode)
	root.Left.Val = 2
	root.Left.Left = new(TreeNode)
	root.Left.Left.Val = 3
	root.Left.Left.Left = new(TreeNode)
	root.Left.Left.Left.Val = 4
	root.Left.Left.Left.Left = new(TreeNode)
	root.Left.Left.Left.Left.Val = 5

	if !reflect.DeepEqual(amountOfTime(root, 4), 3) {
		t.Fatalf("Test_Infection4 Case Failed")
	}
}
