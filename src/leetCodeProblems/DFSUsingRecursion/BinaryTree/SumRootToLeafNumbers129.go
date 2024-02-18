package main

/*
- Leetcode - https://leetcode.com/problems/sum-root-to-leaf-numbers
- Time - O(n)
- Space - O(h)
*/
import "strconv"

func sumNumbersUtil(node *TreeNode, pathSum string, out *int) {
	if node == nil {
		return
	}

	pathSum += strconv.Itoa(node.Val)

	if node.Left == nil && node.Right == nil {
		a, _ := strconv.Atoi(pathSum)
		*out += a
		return
	}
	sumNumbersUtil(node.Left, pathSum, out)
	sumNumbersUtil(node.Right, pathSum, out)
}

func sumNumbers(root *TreeNode) int {
	out := 0
	sumNumbersUtil(root, "", &out)
	return out
}
