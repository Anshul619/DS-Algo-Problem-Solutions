package main

/*
- Leetcode - https://leetcode.com/problems/remove-nth-node-from-end-of-list/description/
*/
import (
	"reflect"
	"testing"
)

func TestRemoveNthFromEnd(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		list := new(ListNode)
		list.Val = 1
		list.Next = new(ListNode)
		list.Next.Val = 2
		list.Next.Next = new(ListNode)
		list.Next.Next.Val = 3
		list.Next.Next.Next = new(ListNode)
		list.Next.Next.Next.Val = 4
		list.Next.Next.Next.Next = new(ListNode)
		list.Next.Next.Next.Next.Val = 5

		if !reflect.DeepEqual(getSlice(removeNthFromEnd(list, 2)), []int{1, 2, 3, 5}) {
			t.Fail()
		}
	})

	t.Run("test2", func(t *testing.T) {
		list := new(ListNode)
		list.Val = 1

		if !reflect.DeepEqual(getSlice(removeNthFromEnd(list, 1)), []int{}) {
			t.Fail()
		}
	})

	t.Run("test3", func(t *testing.T) {
		list := new(ListNode)
		list.Val = 1
		list.Next = new(ListNode)
		list.Next.Val = 2

		if !reflect.DeepEqual(getSlice(removeNthFromEnd(list, 1)), []int{1}) {
			t.Fail()
		}
	})

	t.Run("test4", func(t *testing.T) {
		list := new(ListNode)
		list.Val = 1
		list.Next = new(ListNode)
		list.Next.Val = 2

		if !reflect.DeepEqual(getSlice(removeNthFromEnd(list, 2)), []int{2}) {
			t.Fail()
		}
	})
}
