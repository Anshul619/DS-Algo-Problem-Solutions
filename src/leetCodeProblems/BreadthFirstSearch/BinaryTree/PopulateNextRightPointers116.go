package main

/*
- LeetCode - https://leetcode.com/problems/populating-next-right-pointers-in-each-node/description/
- Time - O(n)
- Space - O(1)
*/
type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	leftMost := root

	// traversing left by left
	for leftMost.Left != nil {
		head := leftMost

		// traversing on the same level
		for head != nil {
			// Connect left -> Right
			head.Left.Next = head.Right

			// Connect Right -> Next left (if exists)
			if head.Next != nil {
				head.Right.Next = head.Next.Left
			}

			head = head.Next
		}

		leftMost = leftMost.Left
	}
	return root
}
