package main

import "testing"

func TestHashCycle1(t *testing.T) {
	head := new(ListNode)
	head.Val = 3
	head.Next = new(ListNode)
	head.Next.Val = 2
	head.Next.Next = new(ListNode)
	head.Next.Next.Val = 0
	head.Next.Next.Next = new(ListNode)
	head.Next.Next.Next.Val = -1
	head.Next.Next.Next.Next = head.Next

	if !hasCycle(head) {
		t.Errorf("failed")
	}
}

func TestHashCycle2(t *testing.T) {
	head := new(ListNode)
	head.Val = 1
	head.Next = new(ListNode)
	head.Next.Val = 2
	head.Next.Next = head

	if !hasCycle(head) {
		t.Errorf("failed")
	}
}

func TestHashCycle3(t *testing.T) {
	head := new(ListNode)
	head.Val = 1

	if hasCycle(head) {
		t.Errorf("failed")
	}
}
