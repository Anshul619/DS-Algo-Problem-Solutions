package main

/*
- LeetCode - https://leetcode.com/problems/binary-tree-level-order-traversal/
- Time - O(n)
- Space - O(n)
*/
import "log"

func levelOrderUtil(node *TreeNode, out *[][]int, level int) {
	if node == nil {
		return
	}

	if level >= len(*out) {
		*out = append(*out, []int{})
	}

	(*out)[level] = append((*out)[level], node.Val)

	levelOrderUtil(node.Left, out, level+1)
	levelOrderUtil(node.Right, out, level+1)
}

func levelOrder(root *TreeNode) [][]int {

	out := [][]int{}

	levelOrderUtil(root, &out, 0)
	return out
}

func main() {

	root := new(TreeNode)
	root.Val = 3
	root.Left = new(TreeNode)
	root.Left.Val = 9
	root.Right = new(TreeNode)
	root.Right.Val = 20
	root.Right.Left = new(TreeNode)
	root.Right.Left.Val = 15
	root.Right.Right = new(TreeNode)
	root.Right.Right.Val = 7

	log.Println(levelOrder(root))
}
