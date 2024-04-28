package main

import (
	"reflect"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		list := new(ListNode)
		list.Val = 2
		list.Next = new(ListNode)
		list.Next.Val = 4
		list.Next.Next = new(ListNode)
		list.Next.Next.Val = 3

		list1 := new(ListNode)
		list1.Val = 5
		list1.Next = new(ListNode)
		list1.Next.Val = 6
		list1.Next.Next = new(ListNode)
		list1.Next.Next.Val = 4

		if !reflect.DeepEqual(getSlice(addTwoNumbers(list, list1)), []int{7, 0, 8}) {
			t.Fail()
		}
	})

	t.Run("test2", func(t *testing.T) {
		list := new(ListNode)
		list.Val = 0

		list1 := new(ListNode)
		list1.Val = 0

		if !reflect.DeepEqual(getSlice(addTwoNumbers(list, list1)), []int{0}) {
			t.Fail()
		}
	})

	t.Run("test3", func(t *testing.T) {
		list := new(ListNode)
		list.Val = 9
		list.Next = new(ListNode)
		list.Next.Val = 9
		list.Next.Next = new(ListNode)
		list.Next.Next.Val = 9
		list.Next.Next.Next = new(ListNode)
		list.Next.Next.Next.Val = 9
		list.Next.Next.Next.Next = new(ListNode)
		list.Next.Next.Next.Next.Val = 9
		list.Next.Next.Next.Next.Next = new(ListNode)
		list.Next.Next.Next.Next.Next.Val = 9
		list.Next.Next.Next.Next.Next.Next = new(ListNode)
		list.Next.Next.Next.Next.Next.Next.Val = 9

		list1 := new(ListNode)
		list1.Val = 9
		list1.Next = new(ListNode)
		list1.Next.Val = 9
		list1.Next.Next = new(ListNode)
		list1.Next.Next.Val = 9
		list1.Next.Next.Next = new(ListNode)
		list1.Next.Next.Next.Val = 9

		if !reflect.DeepEqual(getSlice(addTwoNumbers(list, list1)), []int{8, 9, 9, 9, 0, 0, 0, 1}) {
			t.Fail()
		}
	})

	t.Run("test4", func(t *testing.T) {
		list := new(ListNode)
		list.Val = 5

		list1 := new(ListNode)
		list1.Val = 5

		if !reflect.DeepEqual(getSlice(addTwoNumbers(list, list1)), []int{0, 1}) {
			t.Fail()
		}
	})

	t.Run("test5", func(t *testing.T) {
		list := new(ListNode)
		list.Val = 3
		list.Next = new(ListNode)
		list.Next.Val = 7

		list1 := new(ListNode)
		list1.Val = 9
		list1.Next = new(ListNode)
		list1.Next.Val = 2

		if !reflect.DeepEqual(getSlice(addTwoNumbers(list, list1)), []int{2, 0, 1}) {
			t.Fail()
		}
	})
}
