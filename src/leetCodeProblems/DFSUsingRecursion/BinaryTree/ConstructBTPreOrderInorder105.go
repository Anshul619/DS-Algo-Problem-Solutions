package main

/*
- Leetcode - https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/description/
- Time - O(n)
- Space - O(n)
*/
func dfs(preorder []int, m map[int]int, preIndex *int, inStart int, inEnd int) *TreeNode {
	if inStart > inEnd || *preIndex >= len(preorder) {
		return nil
	}

	node := new(TreeNode)
	node.Val = preorder[*preIndex]
	*preIndex++

	inIndex := m[node.Val]

	node.Left = dfs(preorder, m, preIndex, inStart, inIndex-1)
	node.Right = dfs(preorder, m, preIndex, inIndex+1, inEnd)
	return node
}
func buildTree(preorder []int, inorder []int) *TreeNode {
	m := make(map[int]int)

	for i, v := range inorder {
		m[v] = i
	}

	preIndex := 0

	return dfs(preorder, m, &preIndex, 0, len(inorder)-1)
}
