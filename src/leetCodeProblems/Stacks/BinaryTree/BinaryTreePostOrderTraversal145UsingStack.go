package main

import (
	"log"
)

/*
- Leetcode - https://leetcode.com/problems/binary-tree-postorder-traversal/description/
- Time - O(n)
- Space - O(n)
*/

func (s stackTreeNode) print() {
	d := append([]*TreeNode{}, s...)

	out := []int{}

	for _, v := range d {
		if v != nil {
			out = append(out, v.Val)
		} else {
			out = append(out, -1)
		}
	}

	log.Println("stack", out)
}
func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	out := []int{}
	toProcess := root

	s := stackTreeNode{}

	for !s.isEmpty() || toProcess != nil {

		for toProcess != nil {

			if toProcess.Right != nil {
				s.push(toProcess.Right)
			}

			s.push(toProcess)
			toProcess = toProcess.Left
		}

		node := s.pop()

		if !s.isEmpty() && node.Right == s.peek() {
			toProcess = s.pop()
			s.push(node)
			continue
		}

		out = append(out, node.Val)
	}

	return out
}

// func main() {
// 	root := new(TreeNode) //Returns pointer to TreeNode object
// 	root.Val = 3
// 	root.Left = new(TreeNode)
// 	root.Left.Val = 5
// 	root.Left.Left = new(TreeNode)
// 	root.Left.Left.Val = 6
// 	root.Left.Right = new(TreeNode)
// 	root.Left.Right.Val = 2
// 	root.Left.Right.Left = new(TreeNode)
// 	root.Left.Right.Left.Val = 7
// 	root.Left.Right.Right = new(TreeNode)
// 	root.Left.Right.Right.Val = 4
// 	root.Left.Left.Right = new(TreeNode)
// 	root.Left.Left.Right.Val = 9

// 	root.Right = new(TreeNode)
// 	root.Right.Val = 1
// 	root.Right.Left = new(TreeNode)
// 	root.Right.Left.Val = 0
// 	root.Right.Right = new(TreeNode)
// 	root.Right.Right.Val = 8

// 	log.Println(postorderTraversal(root))
// }
