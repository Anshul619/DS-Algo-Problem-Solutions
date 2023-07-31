package main

/*
- LeetCode - https://leetcode.com/problems/reverse-odd-levels-of-binary-tree/description/
- Time - O(n)
- Space - O(logn) // recursive call stack
*/
import "log"

func reverseOddLevelUtil(node *TreeNode, reverseNode *TreeNode, level int) {

	if node == nil || reverseNode == nil {
		return
	}

	if level%2 != 0 { // Odd Level
		temp := reverseNode.Val
		reverseNode.Val = node.Val
		node.Val = temp
	}

	level++

	reverseOddLevelUtil(node.Left, reverseNode.Right, level)
	reverseOddLevelUtil(node.Right, reverseNode.Left, level)
}

func reverseOddLevels(root *TreeNode) *TreeNode {
	reverseOddLevelUtil(root.Left, root.Right, 1)
	return root
}

func printInOrder(node *TreeNode) {
	if node == nil {
		return
	}

	log.Println(node.Val)
	printInOrder(node.Left)
	printInOrder(node.Right)
}

// func main() {

// 	root := new(TreeNode) //Returns pointer to TreeNode object
// 	root.Val = 2
// 	root.Left = new(TreeNode)
// 	root.Left.Val = 3
// 	root.Left.Left = new(TreeNode)
// 	root.Left.Left.Val = 8
// 	root.Left.Right = new(TreeNode)
// 	root.Left.Right.Val = 13

// 	root.Right = new(TreeNode)
// 	root.Right.Val = 5
// 	root.Right.Left = new(TreeNode)
// 	root.Right.Left.Val = 21
// 	root.Right.Right = new(TreeNode)
// 	root.Right.Right.Val = 34

// 	newRoot := reverseOddLevels(root)

// 	printInOrder(newRoot)

// }
