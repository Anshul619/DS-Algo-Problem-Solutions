package main

/*
- Leetcode - https://leetcode.com/problems/populating-next-right-pointers-in-each-node-ii/description/
- Time - O(N)
- Space - O(1)
*/

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

// Returns leftmost child of the nodes, at the same level as node n
func getNext(n *Node) *Node {
	t := n.Next

	for t != nil {
		if t.Left != nil {
			return t.Left
		}

		if t.Right != nil {
			return t.Right
		}

		t = t.Next
	}
	return nil
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	if root.Left != nil {
		if root.Right != nil {
			root.Left.Next = root.Right
		} else {
			root.Left.Next = getNext(root)
		}
	}

	if root.Right != nil {
		root.Right.Next = getNext(root)

	}

	connect(root.Right)
	connect(root.Left)

	return root
}
