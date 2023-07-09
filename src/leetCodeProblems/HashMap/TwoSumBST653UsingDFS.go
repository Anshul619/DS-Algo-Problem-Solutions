package main

/*
- Leetcode - https://leetcode.com/problems/two-sum-iv-input-is-a-bst/description/
- Time - O(n)
- Space - O(n)
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTargetUtil(node *TreeNode, m map[int]bool, target int) bool {

	if node == nil {
		return false
	}

	if ok := m[target-node.Val]; ok {
		return true
	}

	m[node.Val] = true

	return findTargetUtil(node.Left, m, target) || findTargetUtil(node.Right, m, target)
}

func findTarget(root *TreeNode, k int) bool {

	m := make(map[int]bool)
	return findTargetUtil(root, m, k)
}

// func main() {
// 	root := new(TreeNode) //Returns pointer to TreeNode object
// 	root.Val = 5
// 	root.Left = new(TreeNode)
// 	root.Left.Val = 3
// 	root.Right = new(TreeNode)
// 	root.Right.Val = 6
// 	root.Left.Left = new(TreeNode)
// 	root.Left.Left.Val = 2
// 	root.Left.Right = new(TreeNode)
// 	root.Left.Right.Val = 4
// 	root.Right.Right = new(TreeNode)
// 	root.Right.Right.Val = 7

// 	//log.Println(findTarget(root, 9))
// 	log.Println(findTarget(root, 28))

// }
