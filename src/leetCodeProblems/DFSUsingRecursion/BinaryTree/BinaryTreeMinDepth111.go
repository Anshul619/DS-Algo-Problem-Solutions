package main

import (
	"log"
	"math"
)

/**
* LeetCode - https://leetcode.com/problems/minimum-depth-of-binary-tree/
- Time - O(n)
- Space - O(1)
*/

func minDepthUtil(node *TreeNode, height int, minHeight *int) {
	if node == nil {
		return
	}

	if node.Left == nil && node.Right == nil {
		if height < *minHeight {
			*minHeight = height
		}
		return
	}

	minDepthUtil(node.Left, height+1, minHeight)
	minDepthUtil(node.Right, height+1, minHeight)
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	out := math.MaxInt

	minDepthUtil(root, 1, &out)
	return out
}

func main() {
	root := new(TreeNode) //Returns pointer to TreeNode object
	root.Val = 3
	root.Left = new(TreeNode)
	root.Left.Val = 9

	root.Right = new(TreeNode)
	root.Right.Val = 20
	root.Right.Left = new(TreeNode)
	root.Right.Left.Val = 15
	root.Right.Right = new(TreeNode)
	root.Right.Right.Val = 7

	log.Println(minDepth(root))
}
