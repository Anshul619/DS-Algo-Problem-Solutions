package main

/*
- Leetcode - https://leetcode.com/problems/flatten-binary-tree-to-linked-list/description/
- Time - O(n)
- Exta Space - O(1)
- Aux Space - O(n)
*/

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	if root.Left != nil {
		flatten(root.Left)
	}

	if root.Right != nil {
		flatten(root.Right)
	}

	if root.Left != nil {
		rootRight := root.Right
		root.Right = root.Left

		t := root.Left

		for t.Right != nil {
			t = t.Right
		}
		t.Right = rootRight
	}

	root.Left = nil
}
