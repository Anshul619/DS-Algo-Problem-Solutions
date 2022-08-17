package main

import "log"

/*
LeetCode - https://leetcode.com/problems/evaluate-boolean-binary-tree/
*/

/*type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}*/

func evaluateTree(root *TreeNode) bool {

	if root.Left == nil && root.Right == nil {
		if root.Val == 0 {
			return false
		} else {
			return true
		}
	}

	if root.Val == 2 {
		return evaluateTree(root.Left) || evaluateTree(root.Right)
	} else {
		return evaluateTree(root.Left) && evaluateTree(root.Right)
	}
}

func main() {

	root := new(TreeNode) //Returns pointer to TreeNode object
	root.Val = 2
	root.Left = new(TreeNode)
	root.Left.Val = 1
	root.Right = new(TreeNode)
	root.Right.Val = 3
	root.Right.Left = new(TreeNode)
	root.Right.Left.Val = 0
	root.Right.Right = new(TreeNode)
	root.Right.Right.Val = 1

	log.Println(evaluateTree(root))

}
