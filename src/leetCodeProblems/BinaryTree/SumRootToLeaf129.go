package main

/*
- LeetCode - https://leetcode.com/problems/sum-root-to-leaf-numbers
*/

import (
	"log"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbersUtil(node *TreeNode, pathVal int) int {

	if node == nil {
		return 0
	}

	pathVal, _ = strconv.Atoi(strconv.Itoa(pathVal) + strconv.Itoa(node.Val))

	if node.Left == nil && node.Right == nil { // Leaf node, hence return the sum
		return pathVal
	}

	return sumNumbersUtil(node.Left, pathVal) + sumNumbersUtil(node.Right, pathVal)
}

func sumNumbers(root *TreeNode) int {
	return sumNumbersUtil(root, 0)
}

func main() {

	root := new(TreeNode) //Returns pointer to TreeNode object
	root.Val = 1
	root.Left = new(TreeNode)
	root.Left.Val = 2
	root.Right = new(TreeNode)
	root.Right.Val = 3

	log.Println(sumNumbers(root))
}
