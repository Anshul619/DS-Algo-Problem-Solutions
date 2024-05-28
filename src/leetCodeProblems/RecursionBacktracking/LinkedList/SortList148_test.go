package main

import (
	"reflect"
	"testing"
)

func TestSortList(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		list := new(ListNode)
		list.Val = 4
		list.Next = new(ListNode)
		list.Next.Val = 2
		list.Next.Next = new(ListNode)
		list.Next.Next.Val = 1
		list.Next.Next.Next = new(ListNode)
		list.Next.Next.Next.Val = 3

		if !reflect.DeepEqual(getSlice(sortList(list)), []int{1, 2, 3, 4}) {
			t.Fail()
		}
	})

	t.Run("test2", func(t *testing.T) {
		head := new(ListNode)
		head.Val = -1
		head.Next = new(ListNode)
		head.Next.Val = 5
		head.Next.Next = new(ListNode)
		head.Next.Next.Val = 3
		head.Next.Next.Next = new(ListNode)
		head.Next.Next.Next.Val = 4
		head.Next.Next.Next.Next = new(ListNode)
		head.Next.Next.Next.Next.Val = 0

		if !reflect.DeepEqual(getSlice(sortList(head)), []int{-1, 0, 3, 4, 5}) {
			t.Fail()
		}
	})

	t.Run("test3", func(t *testing.T) {
		list := new(ListNode)
		list.Val = -1
		list.Next = new(ListNode)
		list.Next.Val = 5

		if !reflect.DeepEqual(getSlice(sortList(list)), []int{-1, 5}) {
			t.Fail()
		}
	})

	t.Run("test4", func(t *testing.T) {
		head := new(ListNode)
		head.Val = 1
		head.Next = new(ListNode)
		head.Next.Val = 3
		head.Next.Next = new(ListNode)
		head.Next.Next.Val = 3
		head.Next.Next.Next = new(ListNode)
		head.Next.Next.Next.Val = 1
		head.Next.Next.Next.Next = new(ListNode)
		head.Next.Next.Next.Next.Val = 3

		if !reflect.DeepEqual(getSlice(sortList(head)), []int{1, 1, 3, 3, 3}) {
			t.Fail()
		}
	})
}
