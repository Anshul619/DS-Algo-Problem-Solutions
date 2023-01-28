package main

/*
- LeetCode - https://leetcode.com/problems/binary-search-tree-to-greater-sum-tree/submissions/
*/
import "log"

func util(node *TreeNode, sumToBeAdded int) int {

	if node == nil {
		return sumToBeAdded
	}

	if node.Left == nil && node.Right == nil {
		node.Val += sumToBeAdded
		return node.Val
	}

	//right :=

	// log.Println(node.Val, right, sumToBeAdded)
	node.Val += util(node.Right, sumToBeAdded)

	//left := util(node.Left, node.Val)
	return util(node.Left, node.Val)
}

func bstToGst(root *TreeNode) *TreeNode {

	util(root, 0)
	return root
}

func printTree(node *TreeNode) {

	if node == nil {
		return
	}

	printTree(node.Left)
	log.Println("Out -", node.Val)
	printTree(node.Right)

}

// func main() {
// 	root := new(TreeNode)
// 	root.Val = 4
// 	root.Left = new(TreeNode)
// 	root.Left.Val = 1
// 	root.Right = new(TreeNode)
// 	root.Right.Val = 6
// 	root.Left.Left = new(TreeNode)
// 	root.Left.Left.Val = 0
// 	root.Left.Right = new(TreeNode)
// 	root.Left.Right.Val = 2
// 	root.Left.Right.Right = new(TreeNode)
// 	root.Left.Right.Right.Val = 3

// 	root.Right.Left = new(TreeNode)
// 	root.Right.Left.Val = 5
// 	root.Right.Right = new(TreeNode)
// 	root.Right.Right.Val = 7
// 	root.Right.Right.Right = new(TreeNode)
// 	root.Right.Right.Right.Val = 8

// 	// root := new(TreeNode)
// 	// root.Val = 3
// 	// root.Left = new(TreeNode)
// 	// root.Left.Val = 2
// 	// root.Right = new(TreeNode)
// 	// root.Right.Val = 4
// 	// root.Left.Left = new(TreeNode)
// 	// root.Left.Left.Val = 1

// 	root = bstToGst(root)

// 	printTree(root)
// }
