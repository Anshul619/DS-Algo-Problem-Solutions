package main

//import "log"

/*
LeetCode - https://leetcode.com/problems/evaluate-boolean-binary-tree/
*/

// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

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

//
