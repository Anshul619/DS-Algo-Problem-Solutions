package main

/*
- Leetcode - https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/description/
- Time - O(n)
- Space - O(n)
*/
func buildTreeRecur(preorder []int, m map[int]int, inStart, inEnd int, preIndex *int) *TreeNode {
	if inStart > inEnd || *preIndex >= len(preorder) {
		return nil
	}

	node := new(TreeNode)
	node.Val = preorder[*preIndex]
	*preIndex++

	if inStart == inEnd {
		return node
	}

	inIndex := m[node.Val]
	node.Left = buildTreeRecur(preorder, m, inStart, inIndex-1, preIndex)
	node.Right = buildTreeRecur(preorder, m, inIndex+1, inEnd, preIndex)
	return node
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	m := make(map[int]int)

	for i, v := range inorder {
		m[v] = i
	}

	preIndex := 0

	return buildTreeRecur(preorder, m, 0, len(inorder), &preIndex)
}
