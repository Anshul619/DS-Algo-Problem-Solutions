package main

/*
- Leetcode - https://leetcode.com/problems/path-sum-ii/
- Time - O(N^2)
- Space - O(H)
*/
func pathSumUtil1(node *TreeNode, targetSum int, path []int, out *[][]int) {
	if node == nil {
		return
	}

	targetSum -= node.Val

	path = append(path, node.Val)

	if node.Left == nil && node.Right == nil {
		if targetSum == 0 {
			*out = append(*out, append([]int{}, path...))
		}

		return
	}

	if node.Left != nil {
		pathSumUtil1(node.Left, targetSum, path, out)
	}

	if node.Right != nil {
		pathSumUtil1(node.Right, targetSum, path, out)
	}
}

func pathSum1(root *TreeNode, targetSum int) [][]int {
	out := [][]int{}
	pathSumUtil1(root, targetSum, []int{}, &out)
	return out
}
