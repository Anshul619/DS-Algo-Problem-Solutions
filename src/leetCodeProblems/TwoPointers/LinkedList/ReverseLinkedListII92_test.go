package main

import (
	"reflect"
	"testing"
)

func TestReverseList(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		head := new(ListNode)
		head.Val = 1
		head.Next = new(ListNode)
		head.Next.Val = 2
		head.Next.Next = new(ListNode)
		head.Next.Next.Val = 3
		head.Next.Next.Next = new(ListNode)
		head.Next.Next.Next.Val = 4
		head.Next.Next.Next.Next = new(ListNode)
		head.Next.Next.Next.Next.Val = 5

		if !reflect.DeepEqual(getSlice(reverseBetween(head, 2, 4)), []int{1, 4, 3, 2, 5}) {
			t.Fail()
		}
	})

	t.Run("test2", func(t *testing.T) {
		head := new(ListNode)
		head.Val = 5

		if !reflect.DeepEqual(getSlice(reverseBetween(head, 1, 1)), []int{5}) {
			t.Fail()
		}
	})

	t.Run("test3", func(t *testing.T) {
		head := new(ListNode)
		head.Val = 3
		head.Next = new(ListNode)
		head.Next.Val = 5

		if !reflect.DeepEqual(getSlice(reverseBetween(head, 1, 2)), []int{5, 3}) {
			t.Fail()
		}
	})

	t.Run("test4", func(t *testing.T) {
		head := new(ListNode)
		head.Val = 3
		head.Next = new(ListNode)
		head.Next.Val = 5

		if !reflect.DeepEqual(getSlice(reverseBetween(head, 1, 1)), []int{3, 5}) {
			t.Fail()
		}
	})
}
