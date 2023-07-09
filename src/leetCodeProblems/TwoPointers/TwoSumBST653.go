package main

/*
- Leetcode - https://leetcode.com/problems/two-sum-iv-input-is-a-bst/description/
- Time - O(n)
- Space - O(n)
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorder(node *TreeNode, list *[]int) {

	if node == nil {
		return
	}

	inorder(node.Left, list)
	*list = append(*list, node.Val)
	inorder(node.Right, list)
}

func findTarget(root *TreeNode, k int) bool {

	list := []int{}
	inorder(root, &list)

	left, right := 0, len(list)-1

	for left < right {
		currentSum := list[left] + list[right]

		if currentSum == k {
			return true
		}

		if k > currentSum {
			left++
		} else {
			right--
		}
	}

	return false
}

// func main() {
// 	// root := new(TreeNode) //Returns pointer to TreeNode object
// 	// root.Val = 5
// 	// root.Left = new(TreeNode)
// 	// root.Left.Val = 3
// 	// root.Right = new(TreeNode)
// 	// root.Right.Val = 6
// 	// root.Left.Left = new(TreeNode)
// 	// root.Left.Left.Val = 2
// 	// root.Left.Right = new(TreeNode)
// 	// root.Left.Right.Val = 4
// 	// root.Right.Right = new(TreeNode)
// 	// root.Right.Right.Val = 7

// 	// log.Println(findTarget(root, 9))
// 	// log.Println(findTarget(root, 28))

// 	root := new(TreeNode) //Returns pointer to TreeNode object
// 	root.Val = 2
// 	root.Left = new(TreeNode)
// 	root.Left.Val = 1
// 	root.Right = new(TreeNode)
// 	root.Right.Val = 3
// 	root.Left.Left = new(TreeNode)
// 	root.Left.Left.Val = 2
// 	root.Left.Right = new(TreeNode)
// 	root.Left.Right.Val = 4
// 	root.Right.Right = new(TreeNode)
// 	root.Right.Right.Val = 7

// 	log.Println(findTarget(root, 1))

// }
