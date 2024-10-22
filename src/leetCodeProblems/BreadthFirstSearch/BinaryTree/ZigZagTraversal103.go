package main

/**
- LeetCode - https://leetcode.com/problems/binary-tree-zigzag-level-order-traversal/description/
- Time - O(n)
- Space - O(n)+O(n) = O(n)
**/

type stack []*TreeNode

func (s *stack) push(n *TreeNode) {
	*s = append(*s, n)
}

func (s *stack) pop() *TreeNode {
	n := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return n
}

func (s stack) isEmpty() bool {
	return len(s) == 0
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	s := new(stack)
	out := [][]int{}

	isLeftToRight := true

	s.push(root)

	for !s.isEmpty() {
		s1 := new(stack)
		t := []int{}

		for !s.isEmpty() {
			n := s.pop()

			if isLeftToRight {
				if n.Left != nil {
					s1.push(n.Left)
				}
				if n.Right != nil {
					s1.push(n.Right)
				}
			} else {
				if n.Right != nil {
					s1.push(n.Right)
				}
				if n.Left != nil {
					s1.push(n.Left)
				}
			}

			t = append(t, n.Val)
		}

		s = s1
		out = append(out, t)
		isLeftToRight = !isLeftToRight
	}
	return out
}
