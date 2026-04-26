package main

/*
- Leetcode - https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/
- Time - O(n)
- Space - O(h) (due to recursion stack, where h is height)
*/

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	// If current node matches either. Return it immediately.
	// Why? Because this node could be the LCA.
	if root == p || root == q {
		return root
	}

	// Recurse left and right
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	// If both sides return non-null, current node is the Lowest Common Ancestor
	if left != nil && right != nil {
		return root
	}

	// Otherwise return the non-null side
	if left != nil {
		return left
	}
	return right
}
