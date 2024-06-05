package main

/**
 * https://leetcode.com/problems/validate-binary-search-tree/
 *
 * @author anshul.agrawal
 *
 */

func isValidBSTUtil(root *TreeNode, lower, upper *int) bool {
	if root == nil {
		return true
	}

	if lower != nil && root.Val <= *lower {
		return false
	}

	if upper != nil && root.Val >= *upper {
		return false
	}

	return isValidBSTUtil(root.Left, lower, &root.Val) && isValidBSTUtil(root.Right, &root.Val, upper)
}

func isValidBST(root *TreeNode) bool {
	return isValidBSTUtil(root, nil, nil)
}
