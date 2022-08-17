package main

/**
- LeetCode - https://leetcode.com/problems/root-equals-sum-of-children/
- TimeComplexity - O(1)
- SpaceComplexity - None
*/
import "log"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func checkTree(root *TreeNode) bool {

	if root.Val == root.Left.Val+root.Right.Val {
		return true
	}
	return false
}

func main() {

	root := new(TreeNode) // pointer to TreeNode
	root.Val = 10
	root.Left = new(TreeNode)
	root.Left.Val = 5
	root.Right = new(TreeNode)
	root.Right.Val = 6

	log.Println(checkTree(root))
}
