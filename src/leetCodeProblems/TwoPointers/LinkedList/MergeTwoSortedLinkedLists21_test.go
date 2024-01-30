package main

import "testing"

func TestMergeTwoSortedLists(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		head := new(ListNode)
		head.Val = 1
		head.Next = new(ListNode)
		head.Next.Val = 2
		head.Next.Next = new(ListNode)
		head.Next.Next.Val = 4

		head1 := new(ListNode)
		head1.Val = 1
		head1.Next = new(ListNode)
		head1.Next.Val = 3
		head1.Next.Next = new(ListNode)
		head1.Next.Next.Val = 4

		out := mergeTwoLists(head, head1)

		if out == nil || out.Val != 1 {
			t.Fail()
		}
	})

	t.Run("test2", func(t *testing.T) {
		var head, head1 *ListNode

		out := mergeTwoLists(head, head1)

		if out != nil {
			t.Fail()
		}
	})

	t.Run("test3", func(t *testing.T) {
		var head *ListNode

		head1 := new(ListNode)
		head1.Val = 0

		out := mergeTwoLists(head, head1)

		if out == nil || out.Val != 0 {
			t.Fail()
		}
	})
}
