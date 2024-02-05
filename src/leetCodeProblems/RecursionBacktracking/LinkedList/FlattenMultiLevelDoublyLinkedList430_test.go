package main

/*
- Leetcode - https://leetcode.com/problems/flatten-a-multilevel-doubly-linked-list/description/
*/
import (
	"log"
	"testing"
)

func printLinkedList(head *Node) {

	for head != nil {
		log.Println(head.Val)
		if head.Prev != nil {
			log.Println("Prev", head.Prev.Val)
		}
		if head.Child != nil {
			log.Println("Child", head.Child.Val)
		}
		head = head.Next
	}
}

func TestFlatten(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		head := new(Node)
		head.Val = 8
		head.Next = new(Node)
		head.Next.Val = 9
		head.Next.Prev = head
		head.Next.Next = new(Node)
		head.Next.Next.Val = 10
		head.Next.Next.Prev = head.Next

		child := new(Node)
		child.Val = 11
		child.Next = new(Node)
		child.Next.Val = 12
		child.Next.Prev = child
		head.Child = child
		out := flatten(head)

		printLinkedList(out)
	})
}
