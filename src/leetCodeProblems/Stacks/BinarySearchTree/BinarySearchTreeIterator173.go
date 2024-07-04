package main

/*
- Leetcode - https://leetcode.com/problems/binary-search-tree-iterator/description/
- Time - O(1)
- Space - O(h)
*/
type stack2 []*TreeNode

func (s *stack2) push(t *TreeNode) {
	*s = append(*s, t)
}

func (s *stack2) pop() *TreeNode {
	t := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return t
}

func (s stack2) isEmpty() bool {
	return len(s) == 0
}

type BSTIterator struct {
	s *stack2
}

func pushAllLeftNodes(node *TreeNode, s *stack2) {
	for node != nil {
		s.push(node)
		node = node.Left
	}
}

func Constructor(root *TreeNode) BSTIterator {
	bst := BSTIterator{}
	bst.s = new(stack2)

	pushAllLeftNodes(root, bst.s)
	return bst
}

func (this *BSTIterator) Next() int {
	node := this.s.pop()

	if node.Right != nil {
		pushAllLeftNodes(node.Right, this.s)
	}
	return node.Val
}

func (this *BSTIterator) HasNext() bool {
	return !this.s.isEmpty()
}
