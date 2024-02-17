package main

/*
- LeetCode - https://leetcode.com/problems/binary-tree-right-side-view/description/
- Time - O(h*w)
- Space - O(w)
*/
// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

type queue2 []*TreeNode

func (q *queue2) push(n *TreeNode) {
	*q = append(*q, n)
}

func (q *queue2) pop() *TreeNode {
	t := (*q)[0]
	*q = (*q)[1:]
	return t
}

func (q queue2) isEmpty() bool {
	return len(q) == 0
}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	out := []int{}

	q := new(queue2)
	q.push(root)

	isInserted := false
	for !q.isEmpty() {

		q1 := new(queue2)
		isInserted = false

		for !q.isEmpty() {
			t := q.pop()

			if !isInserted {
				out = append(out, t.Val)
				isInserted = true
			}

			if t.Right != nil {
				q1.push(t.Right)
			}

			if t.Left != nil {
				q1.push(t.Left)
			}
		}

		for !q1.isEmpty() {
			q.push(q1.pop())
		}
	}

	return out
}
