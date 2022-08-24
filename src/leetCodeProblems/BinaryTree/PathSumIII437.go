package main

/*
- LeetCode - https://leetcode.com/problems/path-sum-iii/
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSumUtil(node *TreeNode, pathSum int, targetSum int, m map[int]int) int {
	if node == nil {
		return -1
	}

	pathSum += node.Val

	prefixSum := pathSum - targetSum

	if v, ok := m[prefixSum]; ok {
		return v
	}

	if v, ok := m[pathSum]; ok {
		m[pathSum] = v + 1
	} else {
		m[pathSum] = 1
	}

	pathSumUtil(node.Left, pathSum, targetSum, m)
	pathSumUtil(node.Right, pathSum, targetSum, m)

	if v, ok := m[pathSum]; ok {
		m[pathSum] = v - 1
	}

	return 0
}

func pathSum(root *TreeNode, targetSum int) int {

}
