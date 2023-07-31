package main

/*
- Leetcode - https://leetcode.com/problems/binary-tree-postorder-traversal/description/
- Time - O(n)
- Space - O(1)
*/
func postorderTraversalUtil(root *TreeNode, out *[]int) {
	if root == nil {
		return
	}

	postorderTraversalUtil(root.Left, out)
	postorderTraversalUtil(root.Right, out)
	*out = append(*out, root.Val)
}

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	out := []int{}
	postorderTraversalUtil(root, &out)
	return out
}
