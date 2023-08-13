package main

/*
- LeetCode - https://leetcode.com/problems/balanced-binary-tree
*/
//import "log"

// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

func isBalancedUtil(root *TreeNode) (int, bool) {

	if root == nil {
		return -1, true
	}

	leftHeight, leftIsBalanced := isBalancedUtil(root.Left)
	rightHeight, rightIsBalanced := isBalancedUtil(root.Right)

	leftHeight++
	rightHeight++

	diff, isBalanced, height := leftHeight-rightHeight, false, rightHeight

	if -1 <= diff && diff <= 1 {
		isBalanced = leftIsBalanced && rightIsBalanced
	}

	if leftHeight > rightHeight {
		height = leftHeight
	}

	return height, isBalanced
}

func isBalanced(root *TreeNode) bool {
	_, isBalanced := isBalancedUtil(root)
	return isBalanced
}

// func main() {
// 	root := new(TreeNode) //Returns pointer to TreeNode object
// 	root.Val = 2
// 	root.Left = new(TreeNode)
// 	root.Left.Val = 1
// 	root.Right = new(TreeNode)
// 	root.Right.Val = 3
// 	root.Right.Left = new(TreeNode)
// 	root.Right.Left.Val = 0
// 	root.Right.Right = new(TreeNode)
// 	root.Right.Right.Val = 1
// 	root.Right.Right.Left = new(TreeNode)
// 	root.Right.Right.Left.Val = 4
// 	root.Right.Right.Left.Left = new(TreeNode)
// 	root.Right.Right.Left.Left.Val = 5

// 	//root := new(TreeNode) //Returns pointer to TreeNode object
// 	//root.Val = 1
// 	//root.Left = new(TreeNode)
// 	//root.Left.Val = 2
// 	//root.Left.Left = new(TreeNode)
// 	//root.Left.Left.Val = 3
// 	//root.Left.Left.Left = new(TreeNode)
// 	//root.Left.Left.Left.Val = 4
// 	//root.Right = new(TreeNode)
// 	//root.Right.Val = 2
// 	//root.Right.Right = new(TreeNode)
// 	//root.Right.Right.Val = 3
// 	//root.Right.Right.Right = new(TreeNode)
// 	//root.Right.Right.Right.Val = 4

// 	log.Println(isBalanced(root))

// }
