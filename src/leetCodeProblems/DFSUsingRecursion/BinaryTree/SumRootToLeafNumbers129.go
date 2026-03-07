package main

/*
- Leetcode - https://leetcode.com/problems/sum-root-to-leaf-numbers
- Time - O(n)
- Space - O(h)
*/
func dfs1(root *TreeNode, cur int) int {
	if root == nil {
		return 0
	}

	cur = cur*10 + root.Val

	if root.Left == nil && root.Right == nil {
		return cur
	}
	return dfs1(root.Left, cur) + dfs1(root.Right, cur)
}

func sumNumbers(root *TreeNode) int {
	return dfs1(root, 0)
}
