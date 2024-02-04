package main

/*
- LeetCode - https://leetcode.com/problems/count-good-nodes-in-binary-tree/description/
- Time - O(n)
- Space - O(1)
*/
import (
	"math"
)

func goodNodesUtil(node *TreeNode, maxValInPath int, out *int) {

	if node == nil {
		return
	}

	if node.Val >= maxValInPath {
		*out++
		maxValInPath = node.Val
	}

	goodNodesUtil(node.Left, maxValInPath, out)
	goodNodesUtil(node.Right, maxValInPath, out)
}

func goodNodes(root *TreeNode) int {
	out := new(int)
	goodNodesUtil(root, math.MinInt, out)
	return *out
}

// func main() {

// 	root := new(TreeNode)
// 	root.Val = 3
// 	root.Left = new(TreeNode)
// 	root.Left.Val = 1
// 	root.Left.Left = new(TreeNode)
// 	root.Left.Left.Val = 3

// 	root.Right = new(TreeNode)
// 	root.Right.Val = 4
// 	root.Right.Left = new(TreeNode)
// 	root.Right.Left.Val = 1
// 	root.Right.Right = new(TreeNode)
// 	root.Right.Right.Val = 5

// 	log.Println(goodNodes(root))
// }
