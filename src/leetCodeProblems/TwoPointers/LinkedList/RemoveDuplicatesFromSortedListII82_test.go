package main

/*
- Leetcode - https://leetcode.com/problems/remove-duplicates-from-sorted-list-ii/description
*/
import (
	"reflect"
	"testing"
)

func TestDeleteDuplicatesII(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		list := new(ListNode)
		list.Val = 1
		list.Next = new(ListNode)
		list.Next.Val = 2
		list.Next.Next = new(ListNode)
		list.Next.Next.Val = 3
		list.Next.Next.Next = new(ListNode)
		list.Next.Next.Next.Val = 3
		list.Next.Next.Next.Next = new(ListNode)
		list.Next.Next.Next.Next.Val = 4
		list.Next.Next.Next.Next.Next = new(ListNode)
		list.Next.Next.Next.Next.Next.Val = 4
		list.Next.Next.Next.Next.Next.Next = new(ListNode)
		list.Next.Next.Next.Next.Next.Next.Val = 5

		if !reflect.DeepEqual(getSlice(deleteDuplicatesII(list)), []int{1, 2, 5}) {
			t.Fail()
		}
	})

	t.Run("test2", func(t *testing.T) {
		list := new(ListNode)
		list.Val = 1
		list.Next = new(ListNode)
		list.Next.Val = 1
		list.Next.Next = new(ListNode)
		list.Next.Next.Val = 1
		list.Next.Next.Next = new(ListNode)
		list.Next.Next.Next.Val = 2
		list.Next.Next.Next.Next = new(ListNode)
		list.Next.Next.Next.Next.Val = 3

		if !reflect.DeepEqual(getSlice(deleteDuplicatesII(list)), []int{2, 3}) {
			t.Fail()
		}
	})

	t.Run("test3", func(t *testing.T) {
		if !reflect.DeepEqual(getSlice(deleteDuplicatesII(nil)), []int{}) {
			t.Fail()
		}
	})

	t.Run("test4", func(t *testing.T) {
		list := new(ListNode)
		list.Val = 0
		list.Next = new(ListNode)
		list.Next.Val = 1
		list.Next.Next = new(ListNode)
		list.Next.Next.Val = 2
		list.Next.Next.Next = new(ListNode)
		list.Next.Next.Next.Val = 2
		list.Next.Next.Next.Next = new(ListNode)
		list.Next.Next.Next.Next.Val = 3
		list.Next.Next.Next.Next.Next = new(ListNode)
		list.Next.Next.Next.Next.Next.Val = 4

		if !reflect.DeepEqual(getSlice(deleteDuplicatesII(list)), []int{0, 1, 3, 4}) {
			t.Fail()
		}
	})

	t.Run("test5", func(t *testing.T) {
		list := new(ListNode)
		list.Val = 1
		list.Next = new(ListNode)
		list.Next.Val = 1
		if !reflect.DeepEqual(getSlice(deleteDuplicatesII(list)), []int{}) {
			t.Fail()
		}
	})

	t.Run("test6", func(t *testing.T) {
		list := new(ListNode)
		list.Val = 1
		list.Next = new(ListNode)
		list.Next.Val = 1
		list.Next.Next = new(ListNode)
		list.Next.Next.Val = 1
		list.Next.Next.Next = new(ListNode)
		list.Next.Next.Next.Val = 2
		list.Next.Next.Next.Next = new(ListNode)
		list.Next.Next.Next.Next.Val = 2
		list.Next.Next.Next.Next.Next = new(ListNode)
		list.Next.Next.Next.Next.Next.Val = 3

		if !reflect.DeepEqual(getSlice(deleteDuplicatesII(list)), []int{3}) {
			t.Fail()
		}
	})
}
