package main

/*
- Leetcode - https://leetcode.com/problems/binary-tree-inorder-traversal/description/
- Time - O(n)
- Space - O(1)
*/
func inorderTraversalUtil(root *TreeNode, out *[]int) {
	if root == nil {
		return
	}

	inorderTraversalUtil(root.Left, out)
	*out = append(*out, root.Val)
	inorderTraversalUtil(root.Right, out)
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	out := []int{}
	inorderTraversalUtil(root, &out)
	return out
}
