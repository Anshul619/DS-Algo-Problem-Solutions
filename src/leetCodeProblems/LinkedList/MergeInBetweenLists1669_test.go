package main

import (
	"log"
	"testing"
)

func getListLength(node *ListNode) int {
	length := 0

	for node != nil {
		length++
		node = node.Next
	}

	return length
}

func TestMergeInBetween1(t *testing.T) {
	head := new(ListNode)
	head.Val = 0
	head.Next = new(ListNode)
	head.Next.Val = 1
	head.Next.Next = new(ListNode)
	head.Next.Next.Val = 2
	head.Next.Next.Next = new(ListNode)
	head.Next.Next.Next.Val = 3
	head.Next.Next.Next.Next = new(ListNode)
	head.Next.Next.Next.Next.Val = 4
	head.Next.Next.Next.Next.Next = new(ListNode)
	head.Next.Next.Next.Next.Next.Val = 5

	head1 := new(ListNode)
	head1.Val = 1000000
	head1.Next = new(ListNode)
	head1.Next.Val = 1000001
	head1.Next.Next = new(ListNode)
	head1.Next.Next.Val = 1000002

	new := mergeInBetween(head, 3, 4, head1)

	if getListLength(new) != 7 {
		t.Fatalf("Failed")
	}
}

func TestMergeInBetween2(t *testing.T) {
	head := new(ListNode)
	head.Val = 0
	head.Next = new(ListNode)
	head.Next.Val = 1
	head.Next.Next = new(ListNode)
	head.Next.Next.Val = 2

	head1 := new(ListNode)
	head1.Val = 1000000
	head1.Next = new(ListNode)
	head1.Next.Val = 1000001
	head1.Next.Next = new(ListNode)
	head1.Next.Next.Val = 1000002
	head1.Next.Next.Next = new(ListNode)
	head1.Next.Next.Next.Val = 1000003

	new := mergeInBetween(head, 1, 1, head1)

	if getListLength(new) != 6 {
		t.Fatalf("Failed")
	}
}

func TestMergeInBetween3(t *testing.T) {
	head := new(ListNode)
	head.Val = 0
	head.Next = new(ListNode)
	head.Next.Val = 3
	head.Next.Next = new(ListNode)
	head.Next.Next.Val = 2
	head.Next.Next.Next = new(ListNode)
	head.Next.Next.Next.Val = 1
	head.Next.Next.Next.Next = new(ListNode)
	head.Next.Next.Next.Next.Val = 4
	head.Next.Next.Next.Next.Next = new(ListNode)
	head.Next.Next.Next.Next.Next.Val = 5

	head1 := new(ListNode)
	head1.Val = 1000000
	head1.Next = new(ListNode)
	head1.Next.Val = 1000001
	head1.Next.Next = new(ListNode)
	head1.Next.Next.Val = 1000002

	new := mergeInBetween(head, 3, 4, head1)

	newIterator := new

	for newIterator != nil {
		log.Println("New", newIterator.Val)
		newIterator = newIterator.Next
	}

	if getListLength(new) != 7 {
		t.Fatalf("Failed")
	}
}
