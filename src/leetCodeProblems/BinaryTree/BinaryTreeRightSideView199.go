package main

/*
- LeetCode - https://leetcode.com/problems/binary-tree-right-side-view/submissions/
*/
import "log"

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

func rightSideViewUtil(root *TreeNode, out *[]int, level int) {

	if root == nil {
		return
	}

	if len(*out) == level {
		*out = append(*out, root.Val)
	}

	rightSideViewUtil(root.Right, out, level+1)
	rightSideViewUtil(root.Left, out, level+1)
}

func rightSideView(root *TreeNode) []int {

	var out []int

	rightSideViewUtil(root, &out, 0)

	return out
}

func main() {
	root := new(TreeNode) //Returns pointer to TreeNode object
	root.Val = 2
	root.Left = new(TreeNode)
	root.Left.Val = 1
	root.Right = new(TreeNode)
	root.Right.Val = 3
	root.Right.Left = new(TreeNode)
	root.Right.Left.Val = 0
	root.Right.Right = new(TreeNode)
	root.Right.Right.Val = 1

	log.Println(rightSideView(root))
}
