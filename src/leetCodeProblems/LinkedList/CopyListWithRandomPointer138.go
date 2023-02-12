package main

/*
- LeetCode - https://leetcode.com/problems/copy-list-with-random-pointer/description/
- Time - O(n)
- Space - O(n)
*/
import "log"

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {

	m := make(map[*Node]*Node)

	oldNext := head

	var newHead *Node
	var newNext *Node

	for oldNext != nil {
		newNode := new(Node)
		newNode.Val = oldNext.Val

		m[oldNext] = newNode

		if newHead == nil {
			newHead = newNode
			newNext = newNode
		} else {
			newNext.Next = newNode
			newNext = newNext.Next
		}
		oldNext = oldNext.Next
	}

	oldNext = head

	for oldNext != nil {

		m[oldNext].Random = m[oldNext.Random]
		oldNext = oldNext.Next
	}

	return newHead
}

func main() {
	head := new(Node)
	head.Val = 7
	head.Next = new(Node)

	head.Next.Val = 13
	head.Next.Next = new(Node)
	head.Next.Next.Val = 11
	head.Next.Next.Next = new(Node)
	head.Next.Next.Next.Val = 10
	head.Next.Next.Next.Next = new(Node)
	head.Next.Next.Next.Next.Val = 1

	head.Random = nil
	head.Next.Random = head
	head.Next.Next.Random = head.Next.Next.Next.Next

	newHead := copyRandomList(head)

	next := newHead

	for next != nil {
		log.Println(next.Val)
		log.Println(next.Random)
		next = next.Next
	}
}
