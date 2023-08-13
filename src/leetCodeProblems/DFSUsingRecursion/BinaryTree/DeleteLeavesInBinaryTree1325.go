package main

/*
- LeetCode - https://leetcode.com/problems/delete-leaves-with-a-given-value/solutions/
- Time - O(n)
- Space - O(1)
*/
import "log"

func removeLeafNodes(root *TreeNode, target int) *TreeNode {
	if root == nil {
		return nil
	}

	root.Left = removeLeafNodes(root.Left, target)
	root.Right = removeLeafNodes(root.Right, target)

	if root.Val == target && root.Left == nil && root.Right == nil {
		return nil
	}

	return root
}

func preOrder(node *TreeNode) {

	if node == nil {
		return
	}

	log.Println(node.Val)
	preOrder(node.Left)
	preOrder(node.Right)
}

// func main() {

// 	// root := new(TreeNode) //Returns pointer to TreeNode object
// 	// root.Val = 1
// 	// root.Left = new(TreeNode)
// 	// root.Left.Val = 2
// 	// root.Right = new(TreeNode)
// 	// root.Right.Val = 3
// 	// root.Left.Left = new(TreeNode)
// 	// root.Left.Left.Val = 2

// 	// root.Right.Left = new(TreeNode)
// 	// root.Right.Left.Val = 2
// 	// root.Right.Right = new(TreeNode)
// 	// root.Right.Right.Val = 4

// 	root := new(TreeNode) //Returns pointer to TreeNode object
// 	root.Val = 1
// 	root.Left = new(TreeNode)
// 	root.Left.Val = 3
// 	root.Right = new(TreeNode)
// 	root.Right.Val = 3
// 	root.Left.Left = new(TreeNode)
// 	root.Left.Left.Val = 3

// 	root.Left.Right = new(TreeNode)
// 	root.Left.Right.Val = 2

// 	// root := new(TreeNode) //Returns pointer to TreeNode object
// 	// root.Val = 1
// 	// root.Left = new(TreeNode)
// 	// root.Left.Val = 1
// 	// root.Right = new(TreeNode)
// 	// root.Right.Val = 1

// 	root = removeLeafNodes(root, 3)

// 	preOrder(root)

// }
