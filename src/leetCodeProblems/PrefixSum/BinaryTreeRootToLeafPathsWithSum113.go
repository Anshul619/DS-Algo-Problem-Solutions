package main

/*
- Leetcode - https://leetcode.com/problems/path-sum-ii/
- Time - O(n)
- Space - O(1)
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

// func main() {
// 	root := new(TreeNode) //Returns pointer to TreeNode object
// 	root.Val = 5
// 	root.Left = new(TreeNode)
// 	root.Left.Val = 4
// 	root.Left.Left = new(TreeNode)
// 	root.Left.Left.Val = 11
// 	root.Left.Left.Left = new(TreeNode)
// 	root.Left.Left.Left.Val = 7
// 	root.Left.Left.Right = new(TreeNode)
// 	root.Left.Left.Right.Val = 2

// 	root.Right = new(TreeNode)
// 	root.Right.Val = 8
// 	root.Right.Left = new(TreeNode)
// 	root.Right.Left.Val = 13
// 	root.Right.Right = new(TreeNode)
// 	root.Right.Right.Val = 4
// 	root.Right.Right.Left = new(TreeNode)
// 	root.Right.Right.Left.Val = 5
// 	root.Right.Right.Right = new(TreeNode)
// 	root.Right.Right.Right.Val = 1

// 	log.Println(pathSum1(root, 22))
// }
