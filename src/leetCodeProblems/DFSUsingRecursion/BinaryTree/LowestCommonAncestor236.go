package main

/*
- Leetcode - https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/description/
- Time - O(n)
- Space - O(1)
*/
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == p.Val || root.Val == q.Val {
		return root
	}

	lcaLeft := lowestCommonAncestor(root.Left, p, q)

	lcaRight := lowestCommonAncestor(root.Right, p, q)

	if lcaLeft != nil && lcaRight != nil {
		return root
	}

	if lcaLeft != nil {
		return lcaLeft
	}
	return lcaRight
}
