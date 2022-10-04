package main

/*
- LeetCode - https://leetcode.com/problems/kth-smallest-element-in-a-bst/
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inOrderTraversal(node *TreeNode, k int) (int, int) {
	if node == nil || k <= 0 {
		return -1, k
	}

	left, k := inOrderTraversal(node.Left, k)

	if left != -1 {
		return left, k
	}

	if k--; k == 0 {
		return node.Val, k
	}

	return inOrderTraversal(node.Right, k)
}

func kthSmallest(root *TreeNode, k int) int {
	out, _ := inOrderTraversal(root, k)
	return out
}

// func main() {

// 	//root := new(TreeNode)
// 	//root.Val = 3
// 	//root.Left = new(TreeNode)
// 	//root.Left.Val = 1
// 	//root.Right = new(TreeNode)
// 	//root.Right.Val = 4
// 	//root.Left.Right = new(TreeNode)
// 	//root.Left.Right.Val = 2

// 	root := new(TreeNode)
// 	root.Val = 5
// 	root.Left = new(TreeNode)
// 	root.Left.Val = 3
// 	root.Right = new(TreeNode)
// 	root.Right.Val = 6
// 	root.Left.Left = new(TreeNode)
// 	root.Left.Left.Val = 2
// 	root.Left.Right = new(TreeNode)
// 	root.Left.Right.Val = 4
// 	root.Left.Left.Left = new(TreeNode)
// 	root.Left.Left.Left.Val = 1

// 	log.Println(kthSmallest(root, 1))
// 	log.Println(kthSmallest(root, 2))
// 	log.Println(kthSmallest(root, 3))
// 	log.Println(kthSmallest(root, 4))
// 	log.Println(kthSmallest(root, 5))
// 	log.Println(kthSmallest(root, 6))
// }
