package main

/*
- Leetcode - https://leetcode.com/problems/same-tree/description/
- Time - O(n)
- Space - O(h)
*/
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil || p.Val != q.Val {
		return false
	}

	return (isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right))
}
