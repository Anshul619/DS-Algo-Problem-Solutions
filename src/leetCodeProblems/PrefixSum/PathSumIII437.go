package main

/*
- LeetCode - https://leetcode.com/problems/path-sum-iii/
- Time - O(n)
- Space - O(n)
*/

func pathSumUtil(node *TreeNode, pathSum, targetSum int, out *int, m map[int]int) {
	if node == nil {
		return
	}

	pathSum += node.Val
	prefixSum := pathSum - targetSum

	if v, ok := m[prefixSum]; ok {
		*out += v
	}

	m[pathSum]++

	pathSumUtil(node.Left, pathSum, targetSum, out, m)
	pathSumUtil(node.Right, pathSum, targetSum, out, m)

	m[pathSum]--
}

func pathSum(root *TreeNode, tSum int) int {

	m := make(map[int]int)
	m[0] = 1

	out := 0
	pathSumUtil(root, 0, tSum, &out, m)

	return out
}

// func main() {
// 	root := new(TreeNode) //Returns pointer to TreeNode object
// 	root.Val = 10
// 	root.Left = new(TreeNode)
// 	root.Left.Val = 5
// 	root.Left.Left = new(TreeNode)
// 	root.Left.Left.Val = 3
// 	root.Left.Right = new(TreeNode)
// 	root.Left.Right.Val = 2
// 	root.Left.Right.Right = new(TreeNode)
// 	root.Left.Right.Right.Val = 1
// 	root.Left.Left.Left = new(TreeNode)
// 	root.Left.Left.Left.Val = 3
// 	root.Left.Left.Right = new(TreeNode)
// 	root.Left.Left.Right.Val = -2

// 	root.Right = new(TreeNode)
// 	root.Right.Val = -3
// 	root.Right.Right = new(TreeNode)
// 	root.Right.Right.Val = 11

// 	log.Println(pathSum(root, 8))

// 	// root := new(TreeNode) //Returns pointer to TreeNode object
// 	// root.Val = 5
// 	// root.Left = new(TreeNode)
// 	// root.Left.Val = 4
// 	// root.Right = new(TreeNode)
// 	// root.Right.Val = 8

// 	// root.Left.Left = new(TreeNode)
// 	// root.Left.Left.Val = 11
// 	// root.Left.Left.Left = new(TreeNode)
// 	// root.Left.Left.Left.Val = 7
// 	// root.Left.Left.Right = new(TreeNode)
// 	// root.Left.Left.Right.Val = 2

// 	// root.Right.Left = new(TreeNode)
// 	// root.Right.Left.Val = 13

// 	// root.Right.Right = new(TreeNode)
// 	// root.Right.Right.Val = 4

// 	// root.Right.Right.Left = new(TreeNode)
// 	// root.Right.Right.Left.Val = 5
// 	// root.Right.Right.Right = new(TreeNode)
// 	// root.Right.Right.Right.Val = 1

// 	//log.Println(pathSum(root, 22))
// }
