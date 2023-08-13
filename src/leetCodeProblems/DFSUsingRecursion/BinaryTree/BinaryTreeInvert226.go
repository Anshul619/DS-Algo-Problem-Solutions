package main

/*
- Leetcode - https://leetcode.com/problems/invert-binary-tree/description/
- Time - O(n)
- Space - O(1)
*/

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	invertedRightChild := invertTree(root.Right)
	invertedLeftChild := invertTree(root.Left)

	root.Left = invertedRightChild
	root.Right = invertedLeftChild

	return root
}

// func main() {
// 	root := new(TreeNode) //Returns pointer to TreeNode object
// 	root.Val = 4
// 	root.Left = new(TreeNode)
// 	root.Left.Val = 2
// 	root.Left.Left = new(TreeNode)
// 	root.Left.Left.Val = 1
// 	root.Left.Right = new(TreeNode)
// 	root.Left.Right.Val = 3

// 	root.Right = new(TreeNode)
// 	root.Right.Val = 7
// 	root.Right.Left = new(TreeNode)
// 	root.Right.Left.Val = 6
// 	root.Right.Right = new(TreeNode)
// 	root.Right.Right.Val = 9

// 	invertTree(root)
// 	printInOrder(root)

// 	// root := new(TreeNode) //Returns pointer to TreeNode object
// 	// root.Val = 1
// 	// root.Left = new(TreeNode)
// 	// root.Left.Val = 2

// 	// invertTree(root)
// 	// printInOrder(root)
// }
