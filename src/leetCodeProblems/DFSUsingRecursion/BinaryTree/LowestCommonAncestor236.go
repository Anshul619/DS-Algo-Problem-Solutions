package main

/*
- Leetcode - https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/description/
- Time - O(n)
- Space - O(h)
*/
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	// If the current node is p or q, return it (this node could be the LCA).
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}

	// Explore both sides.
	leftLCA := lowestCommonAncestor(root.Left, p, q)
	rightLCA := lowestCommonAncestor(root.Right, p, q)

	// Both p and q found in different branches → this node is their LCA.
	if leftLCA != nil && rightLCA != nil {
		return root
	}

	// Either both are in the same subtree, or not found at all.
	if leftLCA != nil {
		return leftLCA
	}

	return rightLCA
}
