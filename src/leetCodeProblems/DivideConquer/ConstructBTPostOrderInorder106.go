package main

/*
- LeetCode - https://leetcode.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/description/
- Time - O(n)
- Space - O(n)
*/

func buildSubtree(inorderMap map[int]int, inStart, inEnd, postStart, postEnd int, postorder []int) *TreeNode {
	// Recursion base case
	if inStart > inEnd {
		return nil
	}

	// Recursion base case
	if postStart > postEnd {
		return nil
	}

	root := &TreeNode{}
	root.Val = postorder[postEnd]

	rootIndexInorder := inorderMap[root.Val]
	inLen := inEnd - rootIndexInorder

	root.Left = buildSubtree(inorderMap, inStart, rootIndexInorder-1, postStart, postEnd-inLen-1, postorder)
	root.Right = buildSubtree(inorderMap, rootIndexInorder+1, inEnd, postEnd-inLen-1, postEnd-1, postorder)
	return root
}

func buildTree1(inorder []int, postorder []int) *TreeNode {
	inorderMap := make(map[int]int)

	for i, v := range inorder {
		inorderMap[v] = i
	}

	return buildSubtree(inorderMap, 0, len(postorder)-1, 0, len(inorder)-1, postorder)
}
