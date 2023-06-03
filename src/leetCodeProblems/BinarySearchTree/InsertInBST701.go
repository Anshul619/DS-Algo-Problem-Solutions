package main

/*
- LeetCode - https://leetcode.com/problems/insert-into-a-binary-search-tree/description/
- Time - O(n)
- Space - O(1)
*/
import "log"

func insertUtil(node *TreeNode, parent *TreeNode, isLeftChild bool, val int) {

	if node == nil {
		newNode := new(TreeNode)
		newNode.Val = val

		if isLeftChild {
			parent.Left = newNode
		} else {
			parent.Right = newNode
		}

		return
	}

	if node.Val > val {
		insertUtil(node.Left, node, true, val)
	} else {
		insertUtil(node.Right, node, false, val)
	}
}

func insertIntoBST(root *TreeNode, val int) *TreeNode {

	if root == nil {
		root = new(TreeNode)
		root.Val = val
	} else {
		insertUtil(root, nil, true, val)
	}

	return root
}

func inorder1(node *TreeNode) {

	if node == nil {
		return
	}

	inorder1(node.Left)
	log.Println(node.Val)
	inorder1(node.Right)
}

// func main() {
// 	// root := new(TreeNode)
// 	// root.Val = 4
// 	// root.Left = new(TreeNode)
// 	// root.Left.Val = 2
// 	// root.Left.Left = new(TreeNode)
// 	// root.Left.Left.Val = 1
// 	// root.Left.Right = new(TreeNode)
// 	// root.Left.Right.Val = 3

// 	// root.Right = new(TreeNode)
// 	// root.Right.Val = 7

// 	// newRoot := insertIntoBST(root, 5)

// 	newRoot := insertIntoBST(nil, 5)

// 	inorder1(newRoot)
// }
