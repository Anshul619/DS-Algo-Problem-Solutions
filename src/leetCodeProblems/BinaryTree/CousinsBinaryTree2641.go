package main

/*
- LeetCode - https://leetcode.com/problems/cousins-in-binary-tree-ii/description/
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

func calculateCousinsSum(node *TreeNode, depth int, m map[int]int) {

	if node == nil {
		return
	}

	if depth > 1 {
		m[depth] += node.Val
	}

	calculateCousinsSum(node.Left, depth+1, m)
	calculateCousinsSum(node.Right, depth+1, m)
}

func replaceValueInTreeUtil(node *TreeNode, depth int, m map[int]int) {

	if node == nil {
		return
	}

	if depth == 0 {
		node.Val = 0

		if node.Left != nil {
			node.Left.Val = 0
		}

		if node.Right != nil {
			node.Right.Val = 0
		}
	} else {

		siblingsSum := 0

		if node.Left != nil {
			siblingsSum += node.Left.Val
		}

		if node.Right != nil {
			siblingsSum += node.Right.Val
		}

		cousinsSum := m[depth+1] - siblingsSum

		if node.Left != nil {
			node.Left.Val = cousinsSum
		}

		if node.Right != nil {
			node.Right.Val = cousinsSum
		}
	}

	depth++

	replaceValueInTreeUtil(node.Left, depth, m)
	replaceValueInTreeUtil(node.Right, depth, m)
}

func replaceValueInTree(root *TreeNode) *TreeNode {
	m := make(map[int]int)

	calculateCousinsSum(root, 0, m)
	replaceValueInTreeUtil(root, 0, m)

	return root
}

func main() {

	root := new(TreeNode) //Returns pointer to TreeNode object
	root.Val = 5
	root.Left = new(TreeNode)
	root.Left.Val = 4
	root.Right = new(TreeNode)
	root.Right.Val = 9
	root.Left.Left = new(TreeNode)
	root.Left.Left.Val = 1
	root.Left.Right = new(TreeNode)
	root.Left.Right.Val = 10
	root.Right.Right = new(TreeNode)
	root.Right.Right.Val = 7

	replaceValueInTree(root)

	printInOrder(root)
}
