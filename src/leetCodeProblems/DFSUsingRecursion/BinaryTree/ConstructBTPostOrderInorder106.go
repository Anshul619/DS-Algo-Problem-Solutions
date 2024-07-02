package main

/*
- LeetCode - https://leetcode.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/description/
- Time - O(n)
- Space - O(n)
*/

func buildRecur(postOrder []int, m map[int]int, inStart, inEnd, postStart, postEnd int) *TreeNode {

	if inStart > inEnd {
		return nil
	}

	if postStart > postEnd {
		return nil
	}
	node := new(TreeNode)
	node.Val = postOrder[postEnd]

	inIndex := m[node.Val]

	inLeftLen := inIndex - inStart
	node.Left = buildRecur(postOrder, m, inStart, inIndex-1, postStart, postStart+inLeftLen-1)

	inRightLen := inEnd - inIndex
	node.Right = buildRecur(postOrder, m, inIndex+1, inEnd, postEnd-inRightLen, postEnd-1)
	return node
}
func buildTree1(inorder []int, postorder []int) *TreeNode {
	m := make(map[int]int)

	for i, v := range inorder {
		m[v] = i
	}

	return buildRecur(postorder, m, 0, len(inorder)-1, 0, len(postorder)-1)
}
