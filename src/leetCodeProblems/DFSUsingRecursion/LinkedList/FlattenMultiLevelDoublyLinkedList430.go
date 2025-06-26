package main

/*
- Leetcode - https://leetcode.com/problems/flatten-a-multilevel-doubly-linked-list/description/
- Time - O(n)
- Space - O(1)
*/
type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

// Insert nodeToInsert between root and its next. And return the tail/last node of the list
func insertNode(root, nodeToInsert *Node) *Node {
	nodeToInsertTail := nodeToInsert

	for nodeToInsertTail.Next != nil {
		nodeToInsertTail = nodeToInsertTail.Next
	}

	rootNext := root.Next

	root.Next = nodeToInsert
	nodeToInsert.Prev = root // Doubly list

	nodeToInsertTail.Next = rootNext
	if rootNext != nil {
		rootNext.Prev = nodeToInsertTail // Doubly list
	}

	return rootNext
}

func flatten(root *Node) *Node {
	next := root

	for next != nil {
		if next.Child != nil {
			t := next // Save next as its going to be changed
			next = insertNode(next, flatten(next.Child))
			t.Child = nil
		} else {
			next = next.Next
		}
	}
	return root
}
