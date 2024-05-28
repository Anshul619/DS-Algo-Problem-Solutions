package main

/*
- Leetcode - https://leetcode.com/problems/binary-tree-level-order-traversal/description
*/
type tqueue []*TreeNode

func (q *tqueue) push(n *TreeNode) {
	*q = append(*q, n)
}

func (q *tqueue) pop() *TreeNode {
	n := (*q)[0]
	*q = (*q)[1:]
	return n
}

func (q tqueue) isEmpty() bool {
	return len(q) == 0
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	out := [][]int{}

	q := new(tqueue)
	q.push(root)

	for !q.isEmpty() {
		q1 := new(tqueue)
		l := []int{}

		for !q.isEmpty() {
			e := q.pop()
			if e.Left != nil {
				q1.push(e.Left)
			}
			if e.Right != nil {
				q1.push(e.Right)
			}
			l = append(l, e.Val)
		}

		out = append(out, l)
		q = q1
	}
	return out
}
