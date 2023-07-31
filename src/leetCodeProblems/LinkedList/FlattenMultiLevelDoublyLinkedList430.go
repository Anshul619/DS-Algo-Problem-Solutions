package main

/*
- Leetcode - https://leetcode.com/problems/flatten-a-multilevel-doubly-linked-list/description/
- Time - O(n)
- Space - O(1)
*/
import "log"

type Node1 struct {
	Val   int
	Prev  *Node1
	Next  *Node1
	Child *Node1
}

func insertAtEnd(node *Node1, nodeToInsert *Node1) {
	if nodeToInsert == nil {
		return
	}

	next := node
	for next.Next != nil {
		next = next.Next
	}

	next.Next = nodeToInsert
	nodeToInsert.Prev = next
}
func flatten(root *Node1) *Node1 {
	if root == nil {
		return nil
	}

	next := root

	for next != nil {
		if next.Child != nil {
			temp := next.Next
			next.Next = flatten(next.Child)
			next.Next.Prev = next
			next.Child = nil
			insertAtEnd(next.Next, temp)
		}
		next = next.Next
	}
	return root
}

func printLinkedList1(node *Node1) {

	next := node
	for next != nil {
		log.Println(next.Val)
		if next.Prev != nil {
			log.Println("Prev", next.Prev.Val)
		}
		next = next.Next
	}
}

// func main() {
// 	head := new(Node1)
// 	head.Val = 8
// 	head.Next = new(Node1)
// 	head.Next.Val = 9
// 	head.Next.Prev = head
// 	head.Next.Next = new(Node1)
// 	head.Next.Next.Val = 10
// 	head.Next.Next.Prev = head.Next

// 	child := new(Node1)
// 	child.Val = 11
// 	child.Next = new(Node1)
// 	child.Next.Val = 12
// 	head.Child = child
// 	flatten(head)

// 	printLinkedList1(head)
// }
