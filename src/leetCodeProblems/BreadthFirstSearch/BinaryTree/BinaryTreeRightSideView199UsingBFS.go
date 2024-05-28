package main

/*
- LeetCode - https://leetcode.com/problems/binary-tree-right-side-view/description/
- Time - O(h*w)
- Space - O(w)
*/

type queue2 []*TreeNode

func (q *queue2) push(n *TreeNode) {
	*q = append(*q, n)
}

func (q *queue2) pop() *TreeNode {
	n := (*q)[0]
	*q = (*q)[1:]
	return n
}

func (q queue2) peek() *TreeNode {
	return q[0]
}

func (q queue2) isEmpty() bool {
	return len(q) == 0
}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	q := new(queue2)
	q.push(root)
	out := []int{}

	for !q.isEmpty() {
		q1 := new(queue2)

		out = append(out, q.peek().Val)

		for !q.isEmpty() {
			n := q.pop()

			if n.Right != nil {
				q1.push(n.Right)
			}

			if n.Left != nil {
				q1.push(n.Left)
			}
		}

		q = q1
	}
	return out
}
