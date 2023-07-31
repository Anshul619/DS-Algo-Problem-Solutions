package main

import (
	"log"
)

/*
- Leetcode - https://leetcode.com/problems/binary-tree-postorder-traversal/description/
- Time - O(n)
- Space - O(n)
*/

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	out := []int{}
	current := root

	s := stackTreeNode{}

	for !s.isEmpty() || current != nil {

		node := s.pop()

		if node.Right != nil {
			s.push(node.Right)
		}

		s.push(node)

		if node.Left != nil {
			s.push(node.Left)
		}

		for current != nil {
			s.push(current)
			current = current.Left
		}

		current = s.pop()
		out = append(out, current.Val)

		if current.Right == nil {

			current = nil
			continue
		}

		if !s.isEmpty() && current.Right == s.peek() {
			right := s.peek()
			s.push(current)
			current = right
		}

		current = current.Right
	}

	return out
}

func main() {
	root := new(TreeNode) //Returns pointer to TreeNode object
	root.Val = 3
	root.Left = new(TreeNode)
	root.Left.Val = 5
	root.Left.Left = new(TreeNode)
	root.Left.Left.Val = 6
	root.Left.Right = new(TreeNode)
	root.Left.Right.Val = 2
	root.Left.Right.Left = new(TreeNode)
	root.Left.Right.Left.Val = 7
	root.Left.Right.Right = new(TreeNode)
	root.Left.Right.Right.Val = 4

	root.Right = new(TreeNode)
	root.Right.Val = 1
	root.Right.Left = new(TreeNode)
	root.Right.Left.Val = 0
	root.Right.Right = new(TreeNode)
	root.Right.Right.Val = 8

	log.Println(postorderTraversal(root))
}
