package main

// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

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
