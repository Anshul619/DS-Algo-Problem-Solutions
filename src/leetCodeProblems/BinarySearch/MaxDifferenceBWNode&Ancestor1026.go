package main

/*
- LeetCode - https://leetcode.com/problems/maximum-difference-between-node-and-ancestor/description/
- Time - O(n)
- Space - O(1)
*/
import (
	"log"
	"math"
)

func maxAncestorDiffUtil(node *TreeNode, maxAncestor int, minAncestor int, maxDiff *int) {

	if node == nil {
		return
	}

	// log.Println(node.Val, minAncestor, maxAncestor, *maxDiff)

	if maxAncestor-node.Val > *maxDiff {
		*maxDiff = maxAncestor - node.Val
	} else if node.Val-minAncestor > *maxDiff {
		*maxDiff = node.Val - minAncestor
	}

	if node.Val > maxAncestor {
		maxAncestor = node.Val
	}

	if node.Val < minAncestor {
		minAncestor = node.Val
	}

	maxAncestorDiffUtil(node.Left, maxAncestor, minAncestor, maxDiff)
	maxAncestorDiffUtil(node.Right, maxAncestor, minAncestor, maxDiff)

}

func maxAncestorDiff(root *TreeNode) int {

	out := new(int)
	maxAncestorDiffUtil(root, 0, math.MaxInt, out)

	return *out
}

func main() {

	// head := new(TreeNode)
	// head.Val = 8

	// head.Left = new(TreeNode)
	// head.Left.Val = 3
	// head.Left.Left = new(TreeNode)
	// head.Left.Left.Val = 1
	// head.Left.Right = new(TreeNode)
	// head.Left.Right.Val = 6
	// head.Left.Right.Left = new(TreeNode)
	// head.Left.Right.Left.Val = 4
	// head.Left.Right.Right = new(TreeNode)
	// head.Left.Right.Right.Val = 10

	// head.Right = new(TreeNode)
	// head.Right.Val = 10
	// head.Right.Right = new(TreeNode)
	// head.Right.Right.Val = 14
	// head.Right.Right.Left = new(TreeNode)
	// head.Right.Right.Left.Val = 13

	head := new(TreeNode)
	head.Val = 1

	head.Right = new(TreeNode)
	head.Right.Val = 2
	head.Right.Right = new(TreeNode)
	head.Right.Right.Val = 0
	head.Right.Right.Left = new(TreeNode)
	head.Right.Right.Left.Val = 3

	log.Println(maxAncestorDiff(head))

}
