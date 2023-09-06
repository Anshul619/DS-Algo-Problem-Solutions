package main

/*
- LeetCode - https://leetcode.com/problems/range-sum-of-bst/
- Time - O(logn)
- Space - O(logn) (considering the stack used in recursion)
*/

func rangeSumBSTUtil(root *TreeNode, low, high int, sum *int) {
	if root == nil {
		return
	}

	if low <= root.Val && root.Val <= high {
		*sum += root.Val
	}

	rangeSumBSTUtil(root.Left, low, high, sum)
	rangeSumBSTUtil(root.Right, low, high, sum)
}

func rangeSumBST(root *TreeNode, low int, high int) int {
	sum := 0
	rangeSumBSTUtil(root, low, high, &sum)
	return sum
}

// func main() {
// 	root := new(TreeNode)
// 	root.Val = 4
// 	root.Left = new(TreeNode)
// 	root.Left.Val = 2
// 	root.Right = new(TreeNode)
// 	root.Right.Val = 7
// 	root.Left.Left = new(TreeNode)
// 	root.Left.Left.Val = 1
// 	root.Left.Right = new(TreeNode)
// 	root.Left.Right.Val = 3

// 	//root = searchBST(root, 2)
// 	root = searchBST(root, 5)

// 	log.Println("root", root)
// }
