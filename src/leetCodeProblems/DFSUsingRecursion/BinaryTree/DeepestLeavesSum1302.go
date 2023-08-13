package main

/*
- LeetCode - https://leetcode.com/problems/deepest-leaves-sum
*/

// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

func findLastLevel(node *TreeNode, level int) int {

	if node == nil {
		return level
	}

	level++

	leftLevel := findLastLevel(node.Left, level)
	rightLevel := findLastLevel(node.Right, level)

	if leftLevel < rightLevel {
		return rightLevel
	}

	return leftLevel
}

func deepestLeavesSumUtil(node *TreeNode, reverseLevel int, sum *int) {

	if node == nil {
		return
	}

	reverseLevel--

	if reverseLevel == 0 {
		*sum += node.Val
	}

	deepestLeavesSumUtil(node.Left, reverseLevel, sum)
	deepestLeavesSumUtil(node.Right, reverseLevel, sum)
}

func deepestLeavesSum(root *TreeNode) int {

	maxLevel := findLastLevel(root, 0)
	sum := new(int)

	deepestLeavesSumUtil(root, maxLevel, sum)

	return *sum
}

// func main() {

// 	root := new(TreeNode) //Returns pointer to TreeNode object
// 	root.Val = 1
// 	root.Left = new(TreeNode)
// 	root.Left.Val = 2
// 	root.Right = new(TreeNode)
// 	root.Right.Val = 3
// 	root.Left.Left = new(TreeNode)
// 	root.Left.Left.Val = 4
// 	root.Left.Right = new(TreeNode)
// 	root.Left.Right.Val = 5

// 	root.Left.Left.Left = new(TreeNode)
// 	root.Left.Left.Left.Val = 7

// 	root.Right.Right = new(TreeNode)
// 	root.Right.Right.Val = 6
// 	root.Right.Right.Right = new(TreeNode)
// 	root.Right.Right.Right.Val = 8

// 	log.Println(deepestLeavesSum(root))
// }
