package main

import "testing"

func TestConvertLinkedListToBST(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		head := new(ListNode)
		head.Val = 10
		head.Next = new(ListNode)
		head.Next.Val = -3
		head.Next.Next = new(ListNode)
		head.Next.Next.Val = 0
		head.Next.Next.Next = new(ListNode)
		head.Next.Next.Next.Val = 5
		head.Next.Next.Next.Next = new(ListNode)
		head.Next.Next.Next.Next.Val = 9
		node := sortedListToBST(head)
		if node.Val != 0 {
			t.Fail()
		}
	})
}
