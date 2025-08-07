package main

/*
- Leetcode - https://leetcode.com/problems/remove-nth-node-from-end-of-list/description/
*/
import (
	"reflect"
	"testing"
)

func TestRotateRight(t *testing.T) {
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

		if !reflect.DeepEqual(getSlice(rotateRight(list, 2)), []int{4, 5, 1, 2, 3}) {
			t.Fail()
		}
	})

	t.Run("test2", func(t *testing.T) {
		list := new(ListNode)
		list.Val = 0
		list.Next = new(ListNode)
		list.Next.Val = 1
		list.Next.Next = new(ListNode)
		list.Next.Next.Val = 2

		if !reflect.DeepEqual(getSlice(rotateRight(list, 4)), []int{2, 0, 1}) {
			t.Fail()
		}
	})

	t.Run("test3", func(t *testing.T) {
		if !reflect.DeepEqual(getSlice(rotateRight(nil, 0)), []int{}) {
			t.Fail()
		}
	})

	t.Run("test4", func(t *testing.T) {
		list := new(ListNode)
		list.Val = 1
		list.Next = new(ListNode)
		list.Next.Val = 2

		if !reflect.DeepEqual(getSlice(rotateRight(list, 0)), []int{1, 2}) {
			t.Fail()
		}
	})

	t.Run("test5", func(t *testing.T) {
		list := new(ListNode)
		list.Val = 1
		list.Next = new(ListNode)
		list.Next.Val = 2

		if !reflect.DeepEqual(getSlice(rotateRight(list, 1)), []int{2, 1}) {
			t.Fail()
		}
	})

	t.Run("test6", func(t *testing.T) {
		list := new(ListNode)
		list.Val = 1

		if !reflect.DeepEqual(getSlice(rotateRight(list, 1)), []int{1}) {
			t.Fail()
		}
	})
}
