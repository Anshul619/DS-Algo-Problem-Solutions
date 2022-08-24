package main

import "log"

/*
- LeetCode - https://leetcode.com/problems/path-sum-iii/
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var ans = 0
var m = make(map[int]int)
var targetSum = 0

func pathSumUtil(node *TreeNode, pathSum int) {
	if node == nil {
		return
	}

	pathSum += node.Val
	prefixSum := pathSum - targetSum

	log.Println("-------")
	log.Println("Node Value", node.Val)
	log.Println("Map", m)
	log.Println("pathSum", pathSum)
	log.Println("prefixSum", prefixSum)

	if v, ok := m[prefixSum]; ok {
		log.Println("Path found")
		ans += v
	}

	if v, ok := m[pathSum]; ok {
		m[pathSum] = v + 1
	} else {
		m[pathSum] = 1
	}

	pathSumUtil(node.Left, pathSum)
	pathSumUtil(node.Right, pathSum)

	if v, ok := m[pathSum]; ok {
		m[pathSum] = v - 1
	}
}

func pathSum(root *TreeNode, tSum int) int {

	targetSum = tSum
	ans = 0
	m = make(map[int]int)
	m[0] = 1

	pathSumUtil(root, 0)

	return ans
}

func main() {
	//root := new(TreeNode) //Returns pointer to TreeNode object
	//root.Val = 10
	//root.Left = new(TreeNode)
	//root.Left.Val = 5
	//root.Left.Left = new(TreeNode)
	//root.Left.Left.Val = 3
	//root.Left.Right = new(TreeNode)
	//root.Left.Right.Val = 2
	//root.Left.Right.Right = new(TreeNode)
	//root.Left.Right.Right.Val = 1
	//root.Left.Left.Left = new(TreeNode)
	//root.Left.Left.Left.Val = 3
	//root.Left.Left.Right = new(TreeNode)
	//root.Left.Left.Right.Val = -2
	//
	//root.Right = new(TreeNode)
	//root.Right.Val = -3
	//root.Right.Right = new(TreeNode)
	//root.Right.Right.Val = 11

	root := new(TreeNode) //Returns pointer to TreeNode object
	root.Val = 5
	root.Left = new(TreeNode)
	root.Left.Val = 4
	root.Right = new(TreeNode)
	root.Right.Val = 8

	root.Left.Left = new(TreeNode)
	root.Left.Left.Val = 11
	root.Left.Left.Left = new(TreeNode)
	root.Left.Left.Left.Val = 7
	root.Left.Left.Right = new(TreeNode)
	root.Left.Left.Right.Val = 2

	root.Right.Left = new(TreeNode)
	root.Right.Left.Val = 13

	root.Right.Right = new(TreeNode)
	root.Right.Right.Val = 4

	root.Right.Right.Left = new(TreeNode)
	root.Right.Right.Left.Val = 5
	root.Right.Right.Right = new(TreeNode)
	root.Right.Right.Right.Val = 1

	log.Println(pathSum(root, 22))
}
