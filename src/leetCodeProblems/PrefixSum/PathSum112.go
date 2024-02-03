package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
- Leetcode - https://leetcode.com/problems/path-sum/description/
- Time - O(n)
- Space - O(1)
*/
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	targetSum -= root.Val

	if root.Left == nil && root.Right == nil {
		return targetSum == 0
	}

	if root.Left != nil && hasPathSum(root.Left, targetSum) {
		return true
	}

	if root.Right != nil && hasPathSum(root.Right, targetSum) {
		return true
	}

	return false
}

// func main() {
// 	root := new(TreeNode) //Returns pointer to TreeNode object
// 	root.Val = 5
// 	root.Left = new(TreeNode)
// 	root.Left.Val = 4
// 	root.Left.Left = new(TreeNode)
// 	root.Left.Left.Val = 11
// 	root.Left.Left.Left = new(TreeNode)
// 	root.Left.Left.Left.Val = 7
// 	root.Left.Left.Right = new(TreeNode)
// 	root.Left.Left.Right.Val = 2

// 	root.Right = new(TreeNode)
// 	root.Right.Val = 8
// 	root.Right.Left = new(TreeNode)
// 	root.Right.Left.Val = 13
// 	root.Right.Right = new(TreeNode)
// 	root.Right.Right.Val = 4
// 	root.Right.Right.Right = new(TreeNode)
// 	root.Right.Right.Right.Val = 1

// 	log.Println(hasPathSum(root, 26))
// }
