package main

/**
- LeetCode - https://leetcode.com/problems/binary-tree-zigzag-level-order-traversal/description/
- Time - O(n)
- Space - O(n)+O(n) = O(n)
**/

import "log"

type stack []*TreeNode

func (s *stack) push(e *TreeNode) {
	*s = append(*s, e)
}
func (s *stack) pop() *TreeNode {
	e := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return e
}

func (s stack) isEmpty() bool {
	return len(s) == 0
}

func zigzagLevelOrder(root *TreeNode) [][]int {

	if root == nil {
		return [][]int{}
	}
	currentLevelStack := new(stack)
	nextLevelStack := new(stack)

	currentLevel := 0
	isLeftToRight := true

	out := [][]int{}

	currentLevelStack.push(root)

	for !currentLevelStack.isEmpty() {

		node := currentLevelStack.pop()

		if currentLevel >= len(out) {
			out = append(out, []int{})
		}

		out[currentLevel] = append(out[currentLevel], node.Val)

		if isLeftToRight {
			if node.Left != nil {
				nextLevelStack.push(node.Left)
			}

			if node.Right != nil {
				nextLevelStack.push(node.Right)
			}
		} else {
			if node.Right != nil {
				nextLevelStack.push(node.Right)
			}

			if node.Left != nil {
				nextLevelStack.push(node.Left)
			}
		}

		if currentLevelStack.isEmpty() {
			currentLevel++
			isLeftToRight = !isLeftToRight
			currentLevelStack = nextLevelStack
			nextLevelStack = new(stack)
		}
	}

	return out
}

func main() {
	// s := new(stack)
	// s.push(1)
	// s.push(2)
	// s.push(3)
	// log.Println(s.pop())
	// log.Println(s.pop())

	// root := new(TreeNode) //Returns pointer to TreeNode object
	// root.Val = 1
	// root.Left = new(TreeNode)
	// root.Left.Val = 2
	// root.Right = new(TreeNode)
	// root.Right.Val = 3
	// root.Left.Left = new(TreeNode)
	// root.Left.Left.Val = 4
	// root.Left.Right = new(TreeNode)
	// root.Left.Right.Val = 5

	// root.Left.Left.Left = new(TreeNode)
	// root.Left.Left.Left.Val = 7

	// root.Right.Right = new(TreeNode)
	// root.Right.Right.Val = 6
	// root.Right.Right.Right = new(TreeNode)
	// root.Right.Right.Right.Val = 8

	// root := new(TreeNode) //Returns pointer to TreeNode object
	// root.Val = 3
	// root.Left = new(TreeNode)
	// root.Left.Val = 9
	// root.Right = new(TreeNode)
	// root.Right.Val = 20
	// root.Right.Left = new(TreeNode)
	// root.Right.Left.Val = 15
	// root.Right.Right = new(TreeNode)
	// root.Right.Right.Val = 7

	// root := new(TreeNode) //Returns pointer to TreeNode object
	// root.Val = 1

	log.Println(zigzagLevelOrder(nil))
	//log.Println(zigzagLevelOrder(root))
}
