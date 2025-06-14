package main

/**
 * https://leetcode.com/problems/validate-binary-search-tree/
 * - Time Complexity - O(n)
 * - Space Complexity - O(h)
 *
 * @author anshul.agrawal
 *
 */

// for valid BST, node should be greater than lower and smaller than upper
func isValidBSTRecur(node *TreeNode, lower, upper *int) bool {
	if node == nil {
		return true
	}

	if lower != nil && *lower >= node.Val {
		return false
	}

	if upper != nil && *upper <= node.Val {
		return false
	}

	return isValidBSTRecur(node.Left, lower, &node.Val) && isValidBSTRecur(node.Right, &node.Val, upper)
}

func isValidBST(root *TreeNode) bool {
	return isValidBSTRecur(root, nil, nil)
}
