package main

/*
- LeetCode - https://leetcode.com/problems/populating-next-right-pointers-in-each-node/description/
- Time - O(h * w)
- Space - O(w)
*/
type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

type queue1 []*Node

func (q *queue1) push(n *Node) {
	(*q) = append((*q), n)
}

func (q *queue1) pop() *Node {
	t := (*q)[0]
	(*q) = (*q)[1:]
	return t
}

func (q queue1) isEmpty() bool {
	return len(q) == 0
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	q := new(queue1)
	q.push(root)
	for !q.isEmpty() {

		var prev *Node
		q1 := new(queue1)

		for !q.isEmpty() {
			t := q.pop()
			t.Next = prev
			prev = t

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

	return root
}
