package main

/*
- Leetcode - https://leetcode.com/problems/binary-search-tree-iterator/description/
- Time - O(1)
- Space - O(h)
*/
type stack []*TreeNode

func (s *stack) pop() *TreeNode {
	t := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return t
}

func (s *stack) push(i *TreeNode) {
	*s = append(*s, i)
}

func (s *stack) isEmpty() bool {
	return len(*s) == 0
}

type BSTIterator struct {
	st *stack
}

func (this *BSTIterator) pushLeftNodes(root *TreeNode) {
	for root != nil {
		this.st.push(root)
		root = root.Left
	}
}

func Constructor(root *TreeNode) BSTIterator {
	bst := BSTIterator{}
	bst.st = new(stack)

	bst.pushLeftNodes(root)
	return bst
}

func (this *BSTIterator) Next() int {
	i := this.st.pop()

	this.pushLeftNodes(i.Right)
	return i.Val
}

func (this *BSTIterator) HasNext() bool {
	return !this.st.isEmpty()
}
