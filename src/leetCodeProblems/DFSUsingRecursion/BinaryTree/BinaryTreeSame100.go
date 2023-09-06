package main

import "log"

/*
- Leetcode - https://leetcode.com/problems/same-tree/description/
- Time - O(n)
- Space - O(1)
*/
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil {
		return false
	}

	if q == nil {
		return false
	}

	if p.Val != q.Val {
		return false
	}

	if isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right) {
		return true
	}

	return false
}

func main() {
	root := new(TreeNode) //Returns pointer to TreeNode object
	root.Val = 1
	root.Left = new(TreeNode)
	root.Left.Val = 2
	root.Left.Left = new(TreeNode)
	root.Left.Left.Val = 3
	root.Left.Right = new(TreeNode)
	root.Left.Right.Val = 4

	root.Right = new(TreeNode)
	root.Right.Val = 5
	root.Right.Right = new(TreeNode)
	root.Right.Right.Val = 6

	root1 := new(TreeNode) //Returns pointer to TreeNode object
	root1.Val = 1
	root1.Left = new(TreeNode)
	root1.Left.Val = 2
	root1.Left.Left = new(TreeNode)
	root1.Left.Left.Val = 3
	root1.Left.Right = new(TreeNode)
	root1.Left.Right.Val = 4

	root1.Right = new(TreeNode)
	root1.Right.Val = 5
	root1.Right.Right = new(TreeNode)
	root1.Right.Right.Val = 6

	log.Println(isSameTree(root, root1))
}
